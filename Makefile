TARGET	= light-proxy

all: $(TARGET)
$(TARGET): main.go.orig index.html
	cpp main.go.orig | sed -s "s/^#.*$$//g" > main.go
	go build

.PHONY: clean
clean:
	-rm $(TARGET) main.go

