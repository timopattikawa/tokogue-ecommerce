package com.tokogue.servicemessage.v1.enumeration;

import lombok.Getter;

@Getter
public enum NotificationAuthEmail {

    REGISTRATION_EMAIL("Registration Email Success", "Your email account has been Registered", "REGISTRATION"),
    LOGIN_EMAIL("Login Email Success", "Detection email account login", "LOGIN"),
    PAYMENT_EMAIL("", "", "");

    private String subject;
    private String body;
    private String type;

    NotificationAuthEmail(String subject, String body, String type) {
        this.subject = subject;
        this.body = body;
        this.type = type;
    }

}
