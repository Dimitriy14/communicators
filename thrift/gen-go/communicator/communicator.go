// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package communicator

import (
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// Attributes:
//  - Price
//  - Amount
type Product struct {
  Price float64 `thrift:"price,1" db:"price" json:"price"`
  Amount int32 `thrift:"amount,2" db:"amount" json:"amount"`
}

func NewProduct() *Product {
  return &Product{}
}


func (p *Product) GetPrice() float64 {
  return p.Price
}

func (p *Product) GetAmount() int32 {
  return p.Amount
}
func (p *Product) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.DOUBLE {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Product)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadDouble(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Price = v
}
  return nil
}

func (p *Product)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Amount = v
}
  return nil
}

func (p *Product) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Product"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Product) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("price", thrift.DOUBLE, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:price: ", p), err) }
  if err := oprot.WriteDouble(float64(p.Price)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.price (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:price: ", p), err) }
  return err
}

func (p *Product) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("amount", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:amount: ", p), err) }
  if err := oprot.WriteI32(int32(p.Amount)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.amount (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:amount: ", p), err) }
  return err
}

func (p *Product) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Product(%+v)", *p)
}

type AvgService interface {
  // Parameters:
  //  - Products
  GetAvg(ctx context.Context, products []*Product) (r float64, err error)
}

type AvgServiceClient struct {
  c thrift.TClient
}

func NewAvgServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AvgServiceClient {
  return &AvgServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewAvgServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AvgServiceClient {
  return &AvgServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewAvgServiceClient(c thrift.TClient) *AvgServiceClient {
  return &AvgServiceClient{
    c: c,
  }
}

func (p *AvgServiceClient) Client_() thrift.TClient {
  return p.c
}
// Parameters:
//  - Products
func (p *AvgServiceClient) GetAvg(ctx context.Context, products []*Product) (r float64, err error) {
  var _args0 AvgServiceGetAvgArgs
  _args0.Products = products
  var _result1 AvgServiceGetAvgResult
  if err = p.Client_().Call(ctx, "GetAvg", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

type AvgServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler AvgService
}

func (p *AvgServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *AvgServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *AvgServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewAvgServiceProcessor(handler AvgService) *AvgServiceProcessor {

  self2 := &AvgServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["GetAvg"] = &avgServiceProcessorGetAvg{handler:handler}
return self2
}

func (p *AvgServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x3

}

type avgServiceProcessorGetAvg struct {
  handler AvgService
}

func (p *avgServiceProcessorGetAvg) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := AvgServiceGetAvgArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("GetAvg", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := AvgServiceGetAvgResult{}
var retval float64
  var err2 error
  if retval, err2 = p.handler.GetAvg(ctx, args.Products); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetAvg: " + err2.Error())
    oprot.WriteMessageBegin("GetAvg", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("GetAvg", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Products
type AvgServiceGetAvgArgs struct {
  Products []*Product `thrift:"products,1" db:"products" json:"products"`
}

func NewAvgServiceGetAvgArgs() *AvgServiceGetAvgArgs {
  return &AvgServiceGetAvgArgs{}
}


func (p *AvgServiceGetAvgArgs) GetProducts() []*Product {
  return p.Products
}
func (p *AvgServiceGetAvgArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *AvgServiceGetAvgArgs)  ReadField1(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*Product, 0, size)
  p.Products =  tSlice
  for i := 0; i < size; i ++ {
    _elem4 := &Product{}
    if err := _elem4.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem4), err)
    }
    p.Products = append(p.Products, _elem4)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *AvgServiceGetAvgArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("GetAvg_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AvgServiceGetAvgArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("products", thrift.LIST, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:products: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Products)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.Products {
    if err := v.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
    }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:products: ", p), err) }
  return err
}

func (p *AvgServiceGetAvgArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AvgServiceGetAvgArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AvgServiceGetAvgResult struct {
  Success *float64 `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewAvgServiceGetAvgResult() *AvgServiceGetAvgResult {
  return &AvgServiceGetAvgResult{}
}

var AvgServiceGetAvgResult_Success_DEFAULT float64
func (p *AvgServiceGetAvgResult) GetSuccess() float64 {
  if !p.IsSetSuccess() {
    return AvgServiceGetAvgResult_Success_DEFAULT
  }
return *p.Success
}
func (p *AvgServiceGetAvgResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *AvgServiceGetAvgResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.DOUBLE {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *AvgServiceGetAvgResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadDouble(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *AvgServiceGetAvgResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("GetAvg_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AvgServiceGetAvgResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.DOUBLE, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteDouble(float64(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *AvgServiceGetAvgResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AvgServiceGetAvgResult(%+v)", *p)
}


