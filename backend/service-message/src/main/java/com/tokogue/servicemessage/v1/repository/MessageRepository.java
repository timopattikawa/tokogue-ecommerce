package com.tokogue.servicemessage.v1.repository;

import com.tokogue.servicemessage.v1.domain.Message;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface MessageRepository extends MongoRepository<Message, String> {
}
