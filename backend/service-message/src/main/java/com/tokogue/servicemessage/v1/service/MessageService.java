package com.tokogue.servicemessage.v1.service;

import com.tokogue.servicemessage.commondto.BaseResponse;
import com.tokogue.servicemessage.v1.dto.EmailDTO;
import org.springframework.stereotype.Service;

public interface MessageService {
    public void sendEmailNotification(String message);
}
