package com.tokogue.servicemessage.v1.service;

import com.tokogue.servicemessage.commondto.BaseResponse;
import org.springframework.stereotype.Service;

public interface MessageService {
    public boolean validateOTP(String otp);
}
