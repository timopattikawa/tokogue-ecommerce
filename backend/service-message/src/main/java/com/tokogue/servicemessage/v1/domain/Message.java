package com.tokogue.servicemessage.v1.domain;

import lombok.Data;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.util.Date;

@Data
@Document(collection = "tokogue_member")
public class Message {

    @Id
    private String id;
    @Field(name = "request_type")
    private String requestType;
    @Field(name = "email")
    private String email;
    @Field(name = "created_at")
    private Date createAt;

}
