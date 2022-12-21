Running
-------

Run gatling from maven

```
# Run 150 virtual users, ramp up in 3 seconds for 120 seconds.
mvn gatling:test -Dusers=150 -Dramp=3 -Dduration=120s
```