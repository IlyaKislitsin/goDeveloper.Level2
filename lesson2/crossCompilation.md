# Кросс-компиляция

### под Windows
    $ go build .\main.go
    $ file main.exe
    main.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows

---

### под Linux 
    $ GOOS=linux GOARCH=386 go build main.go
    main: ELF 32-bit LSB executable, Intel 80386, version 1 (SYSV), statically linked, 
    Go BuildID=kBo3K_8h-ltpXFvm33w1/NwNxw3xDCJXZTYgF3du9/9b4467wi9W8DR8fJSzyl/AC04Cbh7Lqrzi1hcZsAi, not stripped

---

### под MacOS
    $ GOOS=darwin GOARCH=amd64 go build main.go
    main: Mach-O 64-bit x86_64 executable
