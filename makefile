.PHONY: rsakey

rsakey:
	rm ./setting/key ./setting/key.pub >> /dev/null; exit 0
	ssh-keygen -t rsa -b 512 -C "git.jediautocare.com" -f ./setting/key -N "" -q