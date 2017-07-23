#Makefile

NAME=/var/www/html/

work:

	vagrant up
	tmux split-window -l 10 'vagrant ssh -- -t "cd ${NAME}/tests/; /bin/bash"'
	tmux split-window -h 'vagrant ssh -- -t "cd ${NAME}; gulp"'
	tmux split-window 'vagrant ssh -- -t "cd ${NAME}; xvfb-run selenium-standalone start"'
	tmux select-pane -t 0
	nvim

