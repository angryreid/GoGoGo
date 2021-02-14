# Package

## Info

1. can be referance in other package(with first name big)
2. one folder, one packge
3. package name can not be same as folder name

## Init method

1. init function exists in package.
2. each package can have mutilple init function
3. init function will execute before main function
4. diffrent init function in diffrent package's order depends on refreance order.

## Remote Package

`go get -u(?) url`

### dependency issue

After go 1.6

### Denpendency Finding

1. searching `verdor` folder in current package
2. searing outside folder until `src/vendor`
3. searching in `GOPATH`
4. searching in `GIROOT`

### Vendor Package manage tik

1. godep (tools/)
2. glide (Masterminds/)
3. dep   (golang/)
