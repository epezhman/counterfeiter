package arguments

const usage = `
USAGE
	counterfeiter
		[-o <output-path>] [-p] [--fake-name <fake-name>]
		[<source-path>] <interface> [-]

ARGUMENTS
	source-path
		Path to the file or directory containing the interface to fake.
		In package mode (-p), source-path should instead specify the path
		of the input package; alternatively you can use the package name
		(e.g. "os") and the path will be inferred from your GOROOT.

	interface
		If source-path is specified: Name of the interface to fake.
		If no source-path is specified: Fully qualified interface path of the interface to fake.
    If -p is specified, this will be the name of the interface to generate.

	example:
		# writes "FakeStdInterface" to ./packagefakes/fake_std_interface.go
		counterfeiter package/subpackage.StdInterface

	'-' argument
		Write code to standard out instead of to a file

OPTIONS
	-o
		Path to the file or directory for the generated fakes.
		This also determines the package name that will be used.
		By default, the generated fakes will be generated in
		the package "xyzfakes" which is nested in package "xyz",
		where "xyz" is the name of referenced package.

	example:
		# writes "FakeMyInterface" to ./mySpecialFakesDir/specialFake.go
		counterfeiter -o ./mySpecialFakesDir/specialFake.go ./mypackage MyInterface

	-p
		Package mode:  When invoked in package mode, counterfeiter
		will generate an interface and shim implementation from a
		package in your GOPATH.  Counterfeiter finds the public methods
		in the package <source-path> and adds those method signatures
		to the generated interface <interface-name>.

	example:
		# generates os.go (interface) and osshim.go (shim) in ${PWD}/osshim
		counterfeiter -p os
		# now generate fake in ${PWD}/osshim/os_fake (fake_os.go)
		go generate osshim/...

	--fake-name
		Name of the fake struct to generate. By default, 'Fake' will
		be prepended to the name of the original interface. (ignored in
		-p mode)

	example:
		# writes "CoolThing" to ./mypackagefakes/cool_thing.go
		counterfeiter --fake-name CoolThing ./mypackage MyInterface
`