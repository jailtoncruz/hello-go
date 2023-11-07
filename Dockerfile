FROM 1.21.3-alpine3.18

COPY . .

CMD [ "go", "run", "." ]