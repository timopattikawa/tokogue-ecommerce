package com.tokogue.servicemessage.util;

import com.tokogue.servicemessage.v1.enumeration.NotificationAuthEmail;
import jakarta.mail.internet.MimeMessage;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.MimeMessageHelper;
import org.springframework.stereotype.Component;


@Component
@Slf4j
@RequiredArgsConstructor
public class EmailUtil {

    @Value("${spring.mail.username}")
    private String email;

    private final JavaMailSender mailSender;

    public void sendEmail(String emailTo, NotificationAuthEmail notificationAuthEmail) {
        log.info("Try send email delete account notification for : {}",
                emailTo);
        try {
            MimeMessage mimeMessage = mailSender.createMimeMessage();
            MimeMessageHelper mimeMessageHelper = new MimeMessageHelper(mimeMessage);

            mimeMessageHelper.setSubject(notificationAuthEmail.getSubject());
            mimeMessageHelper.setFrom("noreply@alfagift.id");
            mimeMessageHelper.setTo(emailTo);
            mimeMessageHelper.setText(notificationAuthEmail.getBody(), true);

            mailSender.send(mimeMessage);
        } catch (Exception e) {
            log.error("Failed send email notification for : {}", emailTo);
            return;
        }
        log.info("Success email notification for : {}", emailTo);
    }

}
