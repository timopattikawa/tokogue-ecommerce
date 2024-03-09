//package com.tokogue.servicemessage.config;
//
//import org.apache.kafka.clients.consumer.ConsumerConfig;
//import org.springframework.context.annotation.Configuration;
//import org.springframework.kafka.annotation.EnableKafka;
//import org.springframework.kafka.core.ConsumerFactory;
//
//import java.util.HashMap;
//import java.util.Map;
//
//@EnableKafka
//@Configuration
//public class KafkaConfig {
//
//    public ConsumerFactory<String, String> consumerFactory() {
//        Map<String, Object> params = new HashMap<>();
//
//        params.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:9092");
//
//        return params;
//    }
//
//}
