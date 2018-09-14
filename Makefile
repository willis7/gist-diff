NAME = gist_diff
PWD := $(MKPATH:%/Makefile=%)


build:
	go build -race -o $(NAME)

run:
	./$(NAME)
