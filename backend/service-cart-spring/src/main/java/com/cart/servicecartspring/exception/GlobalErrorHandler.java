package com.cart.servicecartspring.exception;

import com.cart.servicecartspring.base.BaseErrorResponse;
import com.cart.servicecartspring.base.BaseResponse;
import com.cart.servicecartspring.base.HeaderRequest;
import com.cart.servicecartspring.exception.common.InternalServerError;
import com.cart.servicecartspring.exception.common.NotFoundException;
import org.apache.coyote.BadRequestException;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import java.time.LocalDate;

@ControllerAdvice
public class GlobalErrorHandler extends ResponseEntityExceptionHandler {

    @ExceptionHandler({NotFoundException.class})
    public ResponseEntity<BaseResponse<Object>> handlerNotFound(NotFoundException error,
                                                                HeaderRequest headerRequest) {
        BaseErrorResponse baseErrorResponse = new BaseErrorResponse();
        baseErrorResponse.setMessage(error.getMessage());
        baseErrorResponse.setDate(LocalDate.now());

        return new ResponseEntity<>(BaseResponse.of(
                HttpStatus.NOT_FOUND.value(),
                baseErrorResponse,
                null
        ), HttpStatus.NOT_FOUND);
    }

    @ExceptionHandler({InternalServerError.class})
    public ResponseEntity<BaseResponse<Object>> handlerInternalServerError(InternalServerError error,
                                                                HeaderRequest headerRequest) {
        BaseErrorResponse baseErrorResponse = new BaseErrorResponse();
        baseErrorResponse.setMessage(error.getMessage());
        baseErrorResponse.setDate(LocalDate.now());

        return new ResponseEntity<>(BaseResponse.of(
                HttpStatus.INTERNAL_SERVER_ERROR.value(),
                baseErrorResponse,
                null
        ), HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @ExceptionHandler({BadRequestException.class})
    public ResponseEntity<BaseResponse<Object>> handlerBadRequestException(BadRequestException error,
                                                                           HeaderRequest headerRequest) {
        BaseErrorResponse baseErrorResponse = new BaseErrorResponse();
        baseErrorResponse.setMessage(error.getMessage());
        baseErrorResponse.setDate(LocalDate.now());

        return new ResponseEntity<>(BaseResponse.of(
                HttpStatus.BAD_REQUEST.value(),
                baseErrorResponse,
                null
        ), HttpStatus.BAD_REQUEST);
    }

}
