FROM mongo

COPY init.json /init.json
CMD mongoimport --host mongodb --db eks --collection eks --type json --file /init.json --jsonArray