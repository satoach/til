TARGET  = tmp
OBJS    = main.o
HEADERS =
LIBS    = -lm

CFLAGS = -Wall

all:$(OBJS)
	$(CC) $(CFLAGS) $(OBJS) $(LIBS) -o $(TARGET)

clean:
	-$(RM) $(TARGET)
	-$(RM) *.o

run:$(TARGET)
	./$(TARGET)

.PHONY : all clean run
