package com.tokogue.servicemessage.v1.dto;


import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;
import org.springframework.messaging.handler.annotation.Payload;

@Data
public class EmailDTO {

    @JsonProperty("type")
    String type;

    @JsonProperty("email")
    String email;
}
