pgks := "\
. \
./examples \
"
coverfile := ".coverage"

test *args:
    go test {{ args }} {{ pgks }}

test-cover *args:
    go test {{ args }} -coverprofile .coverage {{ pgks }}

show-coverage:
    go tool cover -html {{ coverfile }}
