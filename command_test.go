package cli

import "testing"

func TestNewCommand(t *testing.T) {
	cmd := NewCommand("testCommand")

	if cmd.Name != "testCommand" {
		t.Errorf("cmd.Name = %v, should have been %v", cmd.Name, "testCommand")
	}
}

func TestAddCommand(t *testing.T) {
	app := NewApp("testApp")
	app.AddCommand("testCommandName")

	if app.Commands["testCommandName"] == nil {
		t.Errorf("app.Commands[%v] is nil", "testCommandName")
	}
}

var testCommand = &Command{
	Name:      "run",
	ShortName: "r",
}

var parseArgsTests = []struct {
	args []string
	ok   bool
}{
	{[]string{"-f"}, true},
	{[]string{"--f"}, true},
	{[]string{"-f", "-p", "80"}, true},
	{[]string{"-f", "--port", "80"}, true},
	{[]string{"run", "-f"}, true},
	{[]string{"run", "--force"}, true},
	{[]string{"run", "-d", "/test"}, true},
	{[]string{"run", "--directory", "/test"}, true},
	{[]string{"run", "-f", "server"}, true},
	{[]string{"run", "-f", "server", " -p", "80"}, true},
	{[]string{"run", "--force", "server", " -p", "80"}, true},
	{[]string{"run", "-f", "server", "--port", "80"}, true},
	{[]string{"run", "-force", "server", "--port", "80"}, true},
	{[]string{"run", "-d", "/test", "server", "-p", "80"}, true},
	{[]string{"run", "--directory", "/test", "server", "--port", "80"}, true},
}

func TestParse(t *testing.T) {
	for _, tt := range parseArgsTests {
		app := NewApp("testapp")
		app.Parse(tt.args)
		// major, minor, ok := ParseHTTPVersion(tt.vers)
		// if ok != tt.ok || major != tt.major || minor != tt.minor {
		// 	type version struct {
		// 		major, minor int
		// 		ok           bool
		// 	}
		// 	t.Errorf("failed to parse %q, expected: %#v, got %#v", tt.vers, version{tt.major, tt.minor, tt.ok}, version{major, minor, ok})
		// }
	}
}
