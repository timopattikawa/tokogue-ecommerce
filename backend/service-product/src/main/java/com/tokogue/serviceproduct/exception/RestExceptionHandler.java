package com.tokogue.serviceproduct.exception;

import com.tokogue.serviceproduct.base.BaseErrorResponse;
import com.tokogue.serviceproduct.base.BaseResponse;
import com.tokogue.serviceproduct.base.GlobalErrorDetail;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

@ControllerAdvice
public class RestExceptionHandler extends ResponseEntityExceptionHandler {

    @ExceptionHandler({NotFoundException.class})
    public ResponseEntity<BaseResponse<Object>> handlerNotFoundException(NotFoundException e) {
        BaseErrorResponse baseErrorResponse = new BaseErrorResponse();
        baseErrorResponse.setErrorCode(4);
        baseErrorResponse.setErrorMessage(e.getMessage());
        baseErrorResponse.setErrorDetail(new GlobalErrorDetail(e.getTrxStatus(), e.getErrorDetail()));
        BaseResponse<Object> response = BaseResponse.of(HttpStatus.NOT_FOUND.value(),
                baseErrorResponse,
                null);
        return new ResponseEntity<>(response, HttpStatus.NOT_FOUND);
    }

    @ExceptionHandler({InternalServerException.class})
    public ResponseEntity<BaseResponse<Object>> handlerInternalServerException(InternalServerException e) {
        BaseErrorResponse baseErrorResponse = new BaseErrorResponse();
        baseErrorResponse.setErrorCode(5);
        baseErrorResponse.setErrorMessage(e.getMessage());
        baseErrorResponse.setErrorDetail(new GlobalErrorDetail(e.getTrxStatus(), e.getErrorDetail()));
        BaseResponse<Object> response = BaseResponse.of(HttpStatus.INTERNAL_SERVER_ERROR.value(),
                baseErrorResponse,
                null);
        return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
    }

}
