# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: qiita/user.proto for package 'QiitaPb'

require 'grpc'
require 'qiita/user_pb'

module QiitaPb
  module UserService
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'qiita.UserService'

      rpc :GetUser, QiitaPb::GetUserRequest, QiitaPb::User
    end

    Stub = Service.rpc_stub_class
  end
end
