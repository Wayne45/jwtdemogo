package main

import io.gatling.javaapi.core.ChainBuilder
import io.gatling.javaapi.core.CoreDsl
import io.gatling.javaapi.core.PopulationBuilder
import io.gatling.javaapi.core.ScenarioBuilder
import org.slf4j.LoggerFactory
import java.time.Duration
import kotlin.time.toJavaDuration

object Configuration {

  private val logger = LoggerFactory.getLogger(Configuration::class.java)

  val host: String by lazy {
    val h = System.getProperty("host")
    h ?: "http://localhost:8081"
  }

  val duration: Duration? by lazy {
    val s = System.getProperty("duration")
    if (s != null) {
      kotlin.time.Duration.parse(s)
        .toJavaDuration()
    } else {
      null
    }
  }
  val users = Integer.getInteger("users", 10)
  val ramp = java.lang.Long.getLong("ramp", 5)

  init {
    logger.info("Duration: $duration, users: $users, ramp: $ramp")
  }
}

fun ScenarioBuilder.configured(chain: ChainBuilder): ScenarioBuilder {
  return if (Configuration.duration != null) {
    this.during(Configuration.duration!!)
      .on(chain)
  } else {
    this.exec(chain)
  }
}

fun ScenarioBuilder.users(): PopulationBuilder {
  return this.injectOpen(CoreDsl.rampUsers(Configuration.users).during(Duration.ofSeconds(Configuration.ramp)))
}
