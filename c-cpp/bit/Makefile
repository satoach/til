TARGET  = target
TESTSH  = test.sh
SRCS=$(wildcard *.c)
OBJS=$(SRCS:.c=.o)



all: clean $(TARGET)

clean:
	@echo $(MAKECMDGOALS)
	rm -rf $(TARGET) *.o

$(TARGET): $(OBJS)
	cc -o $@ $(OBJS) $(LDFLAGS)


.PHONY : all clean test
