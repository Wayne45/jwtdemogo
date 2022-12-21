package simulation

import com.google.gson.Gson
import io.gatling.javaapi.core.Simulation
import io.gatling.javaapi.http.HttpDsl
import main.Configuration

abstract class BaseSimulation : Simulation() {

  val httpProtocol = HttpDsl.http
    .baseUrl(Configuration.host)
    .contentTypeHeader("application/json")

}
