# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: qiita/item.proto for package 'QiitaPb'

require 'grpc'
require 'qiita/item_pb'

module QiitaPb
  module ItemService
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'qiita.ItemService'

      rpc :ListItems, QiitaPb::ListItemsRequest, QiitaPb::ListItemsResponse
      rpc :ListStocks, QiitaPb::ListStocksRequest, QiitaPb::ListStocksResponse
    end

    Stub = Service.rpc_stub_class
  end
end
