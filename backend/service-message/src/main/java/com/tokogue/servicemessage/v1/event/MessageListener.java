package com.tokogue.servicemessage.v1.event;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class MessageListener {

    @KafkaListener(topics = "tokoguemessage", groupId = "message")
    public void getTopicOTP(String message) {
        System.out.println(message);
    }

}
