SRCDIR := src/
INCDIR := inc/
OBJDIR := obj/

CXX := g++
CFLAGS := -std=c++11 -O2 -Wall -I$(INCDIR)
LD := g++

EXE := euler.x
MAIN := euler.cpp

SRCS := $(wildcard $(SRCDIR)/*.cpp)
OBJS := $(addprefix $(OBJDIR)/,$(notdir $(SRCS:.cpp=.o)))

.SUFFIXES:

.SUFFIXES: .o .cpp

$(OBJDIR)/%.o : $(SRCDIR)/%.cpp
	$(CXX) $(CFLAGS) -c -o $@ $<

all : $(EXE)


$(EXE) : $(OBJS) $(MAIN)
	$(LD) $(CFLAGS) -o $@ $^

#$(OBJS) : $(SRCS) $(INCDIR)/*.h | $(OBJDIR)
$(OBJS) : $(SRCS) | $(OBJDIR)

$(OBJDIR):
	mkdir $(OBJDIR)

clean:
	-$(RM) -f $(EXE) $(OBJS)
