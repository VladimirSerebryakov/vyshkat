package console_test

import (
	"bytes"
	"strings"
	"testing"

	"vyshkat/internal/console"
	"vyshkat/internal/inspection"
	"vyshkat/internal/park"
	"vyshkat/internal/reports"
)

// runScript прогоняет сценарий команд через console.Run и возвращает весь вывод.
func runScript(t *testing.T, script string) (string, *park.Park) {
	t.Helper()

	p := park.NewPark()
	reception, err := park.NewReception(inspection.NewServiceCenter(2.0), p)
	if err != nil {
		t.Fatalf("NewReception failed: %v", err)
	}
	svc, err := reports.NewService(p)
	if err != nil {
		t.Fatalf("NewService failed: %v", err)
	}

	var out bytes.Buffer
	if err := console.Run(reception, p, svc, strings.NewReader(script), &out, false); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	return out.String(), p
}

func TestRun_FullScenario(t *testing.T) {
	script := `
add-scooter SC-1 "Самокат Дедлайн" 8 1.0
add-scooter SC-2 "Самокат Энергозатратный" 7 3.5
add-bicycle BC-1 "Велосипед городской" 9 true
add-helmet HM-1 "Шлем спортивный" XL
add-cabinet CC-1 "Зарядный шкаф" 0.5
report-count
report-energy
report-novices
report-inventory
exit
`
	out, p := runScript(t, script)

	cases := []struct {
		testName string
		want     string
	}{
		{"scooter accepted", "принят: Самокат Дедлайн (SC-1)"},
		{"scooter rejected", "отклонён техосмотром: Самокат Энергозатратный (SC-2)"},
		{"bicycle accepted", "принят: Велосипед городской (BC-1)"},
		{"helmet added", "поставлен на баланс: Шлем спортивный (HM-1)"},
		{"count", "Всего единиц в парке: 4"},
		{"energy", "Суммарное суточное энергопотребление: 1.50 кВт·ч"},
		{"novice in list", "  - BC-1: Велосипед городской"},
		{"inventory has item", "  - CC-1: Зарядный шкаф"},
	}
	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			if !strings.Contains(out, tc.want) {
				t.Errorf("output does not contain %q\nfull output:\n%s", tc.want, out)
			}
		})
	}

	if len(p.Vehicles()) != 2 {
		t.Errorf("expected 2 vehicles in park, got %d", len(p.Vehicles()))
	}
}

func TestRun_InvalidInputDoesNotStopLoop(t *testing.T) {
	cases := []struct {
		testName string
		line     string
		want     string
	}{
		{"unknown command", "fly-away", "неизвестная команда"},
		{"wrong args count", "add-helmet HM-1", "ожидается 3 аргумента"},
		{"not a number", "add-scooter SC-1 Самокат eight 1.0", "простота: ожидается целое число"},
		{"bad bool", "add-bicycle BC-1 Велосипед 9 maybe", "сумка: ожидается true или false"},
		{"domain validation", "add-scooter SC-1 Самокат 11 1.0", "simplicity score must be between 1 and 10"},
		{"unclosed quote", `add-helmet HM-1 "Шлем XL`, "незакрытая кавычка"},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			// после ошибочной команды цикл должен продолжиться и выполнить следующую
			out, _ := runScript(t, tc.line+"\nreport-count\nexit\n")
			if !strings.Contains(out, "ошибка: ") || !strings.Contains(out, tc.want) {
				t.Errorf("expected error containing %q\nfull output:\n%s", tc.want, out)
			}
			if !strings.Contains(out, "Всего единиц в парке: 0") {
				t.Errorf("loop stopped after error\nfull output:\n%s", out)
			}
		})
	}
}

func TestRun_DuplicateInventoryNumberReported(t *testing.T) {
	out, p := runScript(t, "add-helmet X-1 Шлем L\nadd-cabinet X-1 Шкаф 0.5\nexit\n")

	if !strings.Contains(out, `inventory number "X-1" is already used`) {
		t.Errorf("expected duplicate number error\nfull output:\n%s", out)
	}
	if len(p.Items()) != 1 {
		t.Errorf("expected 1 item in park, got %d", len(p.Items()))
	}
}

func TestRun_ExitStopsProcessing(t *testing.T) {
	_, p := runScript(t, "exit\nadd-helmet HM-1 Шлем L\n")

	if len(p.Items()) != 0 {
		t.Errorf("commands after exit must not be executed, got %d items", len(p.Items()))
	}
}

func TestRun_EndOfInputWithoutExit(t *testing.T) {
	_, p := runScript(t, "# комментарий игнорируется\nadd-helmet HM-1 Шлем L")

	if len(p.Items()) != 1 {
		t.Errorf("expected 1 item in park, got %d", len(p.Items()))
	}
}
