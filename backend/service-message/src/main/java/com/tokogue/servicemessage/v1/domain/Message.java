package com.tokogue.servicemessage.v1.domain;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.util.Date;

@Document(collection = "tokogue_member")
public class Message {

    @Id
    private String id;
    @Field(name = "otp")
    private int otp;
    @Field(name = "request_type")
    private String requestType;
    @Field(name = "id_user")
    private int idUser;
    @Field(name = "created_at")
    private Date createAt;


}
