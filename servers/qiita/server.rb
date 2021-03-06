require 'bundler'
Bundler.require

lib_dir = File.join(__dir__, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'google/protobuf/well_known_types'
require 'qiita/user_services_pb'
require 'qiita/item_services_pb'

module QiitaClient
  protected

  def qiita_client
    @qiita ||= Qiita::Client.new(access_token: ENV['QIITA_ACCESS_TOKEN'])
  end
end

class Google::Protobuf::Timestamp
  def self.from_time(t)
    Google::Protobuf::Timestamp.new.tap { |pb| pb.from_time(t) }
  end
end

module QiitaPb
  class User
    def self.from_hash(h)
      QiitaPb::User.new(h.slice(*self.descriptor.map(&:name)))
    end
  end

  class Item
    def self.from_hash(h)
      h = h.slice(*self.descriptor.map(&:name))
      h['created_at'] = Google::Protobuf::Timestamp.from_time(Time.parse(h['created_at']))
      h['updated_at'] = Google::Protobuf::Timestamp.from_time(Time.parse(h['updated_at']))
      h['tags'] = h['tags'].map { |t| QiitaPb::Item::Tag.new(t) }
      h['user'] = QiitaPb::User.from_hash(h['user'])
      QiitaPb::Item.new(h)
    end
  end
end

class UserServer < QiitaPb::UserService::Service
  include QiitaClient

  # @param req [QiitaPb::GetUserRequest]
  # @param call [GRPC::ActiveCall]
  # @return [QiitaPb::User]
  def get_user(req, _call)
    resp = qiita_client.get_user(req.user_id)
    QiitaPb::User.from_hash(resp.body)
  end

  # @param req [QiitaPb::ListFollowersRequest]
  # @param call [GRPC::ActiveCall]
  # @return [QiitaPb::ListFollowersRequest]
  def list_followers(req, _call)
    resp = qiita_client.list_user_followers(req.user_id)
    users = resp.body.map { |h| QiitaPb::User.from_hash(h) }
    QiitaPb::ListFollowersResponse.new(users: users)
  end

  # @param req [QiitaPb::ListFolloweesRequest]
  # @param call [GRPC::ActiveCall]
  # @return [QiitaPb::ListFolloweesRequest]
  def list_followees(req, _call)
    resp = qiita_client.list_user_followees(req.user_id)
    users = resp.body.map { |h| QiitaPb::User.from_hash(h) }
    QiitaPb::ListFolloweesResponse.new(users: users)
  end

  # @param req [QiitaPb::ListStockersRequest]
  # @param call [GRPC::ActiveCall]
  # @return [QiitaPb::ListStockersRequest]
  def list_stockers(req, _call)
    resp = qiita_client.list_item_stockers(req.item_id)
    users = resp.body.map { |h| QiitaPb::User.from_hash(h) }
    QiitaPb::ListStockersResponse.new(users: users)
  end
end

class ItemServer < QiitaPb::ItemService::Service
  include QiitaClient

  # @param req [QiitaPb::ListItemsRequest]
  # @param call [GRPC::ActiveCall]
  # @return [QiitaPb::ListItemsResponse]
  def list_items(req, _call)
    resp = qiita_client.list_user_items(req.user_id)
    items = resp.body.map { |h| QiitaPb::Item.from_hash(h) }
    QiitaPb::ListItemsResponse.new(items: items)
  end

  # @param req [QiitaPb::ListStocksRequest]
  # @param call [GRPC::ActiveCall]
  # @return [QiitaPb::ListStocksResponse]
  def list_stocks(req, _call)
    resp = qiita_client.list_user_stocks(req.user_id)
    items = resp.body.map { |h| QiitaPb::Item.from_hash(h) }
    QiitaPb::ListStocksResponse.new(items: items)
  end
end

def main
  s = GRPC::RpcServer.new
  s.add_http2_port('0.0.0.0:50051', :this_port_is_insecure)
  s.handle(UserServer)
  s.handle(ItemServer)
  s.run_till_terminated_or_interrupted([1, 'int', 'SIGQUIT'])
end

main
