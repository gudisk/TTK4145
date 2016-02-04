# Which compiler to use
CC = gcc

# Compiler flags go here.
CFLAGS = -fPIC -std=c11 -Wall -Wextra -g

# Linker flags go here.
LDFLAGS = -shared -lcomedi -lm

# list of sources
ELEVSRC = elev.c io.c test.c

# program executable file name.
TARGET = test.so

# top-level rule, to compile everything.
all: $(TARGET)

# Define all object files.
ELEVOBJ = $(ELEVSRC:.c=.o)

# rule to link the program
$(TARGET): $(ELEVOBJ)
	$(CC) $^ -o $@ $(LDFLAGS)

# Compile: create object files from C source files.
%.o : %.c
	$(CC) $(CFLAGS) -c $< -o $@ 

# rule for cleaning re-compilable files.
clean:
	rm -f $(TARGET) $(ELEVOBJ)

rebuild:	clean all

.PHONY: all rebuild clean
