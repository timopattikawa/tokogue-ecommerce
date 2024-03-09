package com.tokogue.servicemessage.v1.domain;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.util.Date;

@Document(value = "tokogue_message_request_logging")
public class MessageRequestLogging {

    @Id
    private String id;
    @Field("request_type")
    private String requestType;
    @Field("created_at")
    private Date createdAt;
}
