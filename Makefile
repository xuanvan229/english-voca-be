
build:
	docker build -t english-vocab .

up:
	docker-compose up --abort-on-container-exit --remove-orphans englishapp posgres_english

backup-db:
	docker-compose -f docker-compose.yml run --rm -T backup-db | pv > backup.tar.bz2
	# docker-compose -f docker-compose.yml run --rm -T backup-db

restore_db:
	cat backup.tar.bz2 | pv | docker-compose -f docker-compose.yml run --rm -T restore-db


clean:
	docker-compose stop
	docker-compose rm -f

backup:
	pg_dump -d 'english' -U 'postgres' -h 'localhost' -W  > db.sql

restore:
	psql -h 'localhost' -U 'postgres' -d english -1 -f db.sql