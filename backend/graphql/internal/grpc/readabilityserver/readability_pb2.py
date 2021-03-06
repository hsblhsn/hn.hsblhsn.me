# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: readability.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11readability.proto\x12\rreadabilitypb\">\n\x1aGetReadableDocumentRequest\x12\x12\n\nidentifier\x18\x01 \x01(\t\x12\x0c\n\x04html\x18\x02 \x01(\t\":\n\x1bGetReadableDocumentResponse\x12\r\n\x05title\x18\x01 \x01(\t\x12\x0c\n\x04\x62ody\x18\x02 \x01(\t\"-\n\x17GetReadinessInfoRequest\x12\x12\n\nidentifier\x18\x01 \x01(\t\")\n\x18GetReadinessInfoResponse\x12\r\n\x05ready\x18\x01 \x01(\x08\x32\xe4\x01\n\x0bReadability\x12n\n\x13GetReadableDocument\x12).readabilitypb.GetReadableDocumentRequest\x1a*.readabilitypb.GetReadableDocumentResponse\"\x00\x12\x65\n\x10GetReadinessInfo\x12&.readabilitypb.GetReadinessInfoRequest\x1a\'.readabilitypb.GetReadinessInfoResponse\"\x00\x42\x15Z\x13./readabilityclientb\x06proto3')



_GETREADABLEDOCUMENTREQUEST = DESCRIPTOR.message_types_by_name['GetReadableDocumentRequest']
_GETREADABLEDOCUMENTRESPONSE = DESCRIPTOR.message_types_by_name['GetReadableDocumentResponse']
_GETREADINESSINFOREQUEST = DESCRIPTOR.message_types_by_name['GetReadinessInfoRequest']
_GETREADINESSINFORESPONSE = DESCRIPTOR.message_types_by_name['GetReadinessInfoResponse']
GetReadableDocumentRequest = _reflection.GeneratedProtocolMessageType('GetReadableDocumentRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETREADABLEDOCUMENTREQUEST,
  '__module__' : 'readability_pb2'
  # @@protoc_insertion_point(class_scope:readabilitypb.GetReadableDocumentRequest)
  })
_sym_db.RegisterMessage(GetReadableDocumentRequest)

GetReadableDocumentResponse = _reflection.GeneratedProtocolMessageType('GetReadableDocumentResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETREADABLEDOCUMENTRESPONSE,
  '__module__' : 'readability_pb2'
  # @@protoc_insertion_point(class_scope:readabilitypb.GetReadableDocumentResponse)
  })
_sym_db.RegisterMessage(GetReadableDocumentResponse)

GetReadinessInfoRequest = _reflection.GeneratedProtocolMessageType('GetReadinessInfoRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETREADINESSINFOREQUEST,
  '__module__' : 'readability_pb2'
  # @@protoc_insertion_point(class_scope:readabilitypb.GetReadinessInfoRequest)
  })
_sym_db.RegisterMessage(GetReadinessInfoRequest)

GetReadinessInfoResponse = _reflection.GeneratedProtocolMessageType('GetReadinessInfoResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETREADINESSINFORESPONSE,
  '__module__' : 'readability_pb2'
  # @@protoc_insertion_point(class_scope:readabilitypb.GetReadinessInfoResponse)
  })
_sym_db.RegisterMessage(GetReadinessInfoResponse)

_READABILITY = DESCRIPTOR.services_by_name['Readability']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\023./readabilityclient'
  _GETREADABLEDOCUMENTREQUEST._serialized_start=36
  _GETREADABLEDOCUMENTREQUEST._serialized_end=98
  _GETREADABLEDOCUMENTRESPONSE._serialized_start=100
  _GETREADABLEDOCUMENTRESPONSE._serialized_end=158
  _GETREADINESSINFOREQUEST._serialized_start=160
  _GETREADINESSINFOREQUEST._serialized_end=205
  _GETREADINESSINFORESPONSE._serialized_start=207
  _GETREADINESSINFORESPONSE._serialized_end=248
  _READABILITY._serialized_start=251
  _READABILITY._serialized_end=479
# @@protoc_insertion_point(module_scope)
