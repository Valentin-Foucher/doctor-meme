clean:
	rm bin/doctor-meme

run: bin/doctor-meme
	./bin/doctor-meme

save: bin/doctor-meme
	sudo cp ./bin/doctor-meme /usr/local/bin/doctor-meme

bin/doctor-meme:
	go build -o bin/doctor-meme pkg/app/*
