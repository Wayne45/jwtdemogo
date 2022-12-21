package simulation

import io.gatling.javaapi.core.Choice
import io.gatling.javaapi.core.CoreDsl.*
import io.gatling.javaapi.http.HttpDsl.http
import io.gatling.javaapi.http.HttpDsl.status
import main.configured
import main.users
import org.slf4j.LoggerFactory

class BasicSimulation : BaseSimulation() {

    private val logger = LoggerFactory.getLogger(BasicSimulation::class.java)

    val createToken = http("POST /v1/jwt-tokens")
        .post("/v1/jwt-tokens")

    val getToken1 = http("GET  /v1/jwt-tokens")
        .get("/v1/jwt-tokens")
        .header("Authorization", "Bearer #{token}")

    val getToken2 = http("GET  /v1/self")
        .get("/v1/self")
        .header("Authorization", "Bearer #{token}")

    //  val feeder = identityPooledFeeders(10)
    val scn = scenario("Session")
        .configured(
            exec(createToken.check(
                status().shouldBe(200),
                jmesPath("token").saveAs("token")
            ))
                .randomSwitch()
                .on(
                    Choice.withWeight(50.0, exec(getToken1.check(
                        status().shouldBe(200)))),
                    Choice.withWeight(50.0, exec(getToken2.check(
                        status().shouldBe(200))))
                )
        )
//      exec(
//        createToken.check(
//          status().shouldBe(200),
//          jmesPath("sessionToken.token").saveAs("latestToken"))
//      ).exec { session: Session ->
//        val token = session.get<String>("latestToken")
//        logger.debug("latest token: {}", token)
//        TokenFeeder.append(token)
//        val randomToken = TokenFeeder.random()
//        logger.debug("randomToken: {}", randomToken)
//        session.set("token", randomToken)
//      }.repeat(3)
//        .on(exec(
//          getToken.check(
//            status().shouldBe(200))
//        ))
//    )

    init {
        setUp(
            scn.users()
        ).protocols(httpProtocol)
    }

}
