TARGET  = tmp
OBJS    = main.o
HEADERS =
LIBS    =

CFLAGS = -Wall

all:$(OBJS)
	gcc $(CFLAGS) $(OBJS) $(LIBS) -o $(TARGET)

clean:
	-$(RM) $(TARGET)
	-$(RM) *.o

run:$(TARGET)
	./$(TARGET)

.PHONY : all clean run
