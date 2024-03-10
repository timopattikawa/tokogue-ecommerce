package com.tokogue.servicemessage.v1.service.impl;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.tokogue.servicemessage.util.EmailUtil;
import com.tokogue.servicemessage.v1.dto.EmailDTO;
import com.tokogue.servicemessage.v1.enumeration.NotificationAuthEmail;
import com.tokogue.servicemessage.v1.service.MessageService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@Service
@Slf4j
@RequiredArgsConstructor
public class MessageServiceImplementation implements MessageService {

    private final EmailUtil emailUtil;
    private final ObjectMapper mapper;

    @KafkaListener(topics = "tokoguemessage", groupId = "message")
    @Override
    public void sendEmailNotification(String message) {
        log.info("Request Kafka {}", message);
        try {
            EmailDTO emailDTO = mapper.readValue(message, EmailDTO.class);

            if(emailDTO.getType() == null || emailDTO.getType().isEmpty()) {
                log.info("Empty emailDTO from kafka topic to send email");
                return;
            }

            if(emailDTO.getEmail() == null || emailDTO.getEmail().isEmpty()) {
                log.info("Empty email from kafka topic to send email");
                return;
            }

            if(emailDTO.getType().equals(NotificationAuthEmail.REGISTRATION_EMAIL.getType())) {
                emailUtil.sendEmail(emailDTO.getEmail(), NotificationAuthEmail.REGISTRATION_EMAIL);
            } else if(emailDTO.getType().equals(NotificationAuthEmail.LOGIN_EMAIL.getType())) {
                emailUtil.sendEmail(emailDTO.getEmail(), NotificationAuthEmail.LOGIN_EMAIL);
            }
        }catch (Exception e) {
            log.info(e.getMessage());
            throw new RuntimeException();
        }
    }
}
