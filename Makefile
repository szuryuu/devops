.PHONY: dev prod clean build

setup:
	@echo ">> RUNNING ANSIBLE SETUP"
	ansible-playbook nginx/ansible/playbook.yml --ask-become-pass
