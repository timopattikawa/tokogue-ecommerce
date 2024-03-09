package com.tokogue.servicemessage.commondto;

import lombok.AllArgsConstructor;
import lombok.Data;

import java.time.LocalDate;

@Data
public class BaseErrorResponse {
    int status;
    String errorMessage;
    LocalDate date;
}
