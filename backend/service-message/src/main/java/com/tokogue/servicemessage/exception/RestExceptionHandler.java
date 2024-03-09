package com.tokogue.servicemessage.exception;

import com.tokogue.servicemessage.commondto.BaseErrorResponse;
import com.tokogue.servicemessage.commondto.BaseResponse;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import java.time.LocalDate;

@ControllerAdvice
public class RestExceptionHandler extends ResponseEntityExceptionHandler {

    @ExceptionHandler({NotFoundException.class})
    public ResponseEntity<Object> handlerNotFoundException(NotFoundException e) {
        BaseErrorResponse baseErrorResponse = new BaseErrorResponse();
        baseErrorResponse.setStatus(HttpStatus.NOT_FOUND.value());
        baseErrorResponse.setErrorMessage(e.getMessage());
        baseErrorResponse.setDate(LocalDate.now());
        return new ResponseEntity<>(new BaseResponse<Object>(HttpStatus.NOT_FOUND.value(), baseErrorResponse, new Object()),
                HttpStatus.BAD_REQUEST);
    }

    @ExceptionHandler
    public ResponseEntity<Object> handlerUnauthorizedException(UnauthorizedException e) {
        BaseErrorResponse baseErrorResponse = new BaseErrorResponse();
        baseErrorResponse.setStatus(HttpStatus.UNAUTHORIZED.value());
        baseErrorResponse.setErrorMessage(e.getMessage());
        baseErrorResponse.setDate(LocalDate.now());
        return new ResponseEntity<>(baseErrorResponse, HttpStatus.UNAUTHORIZED);
    }

}
