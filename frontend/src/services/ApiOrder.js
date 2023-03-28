import axios from 'axios';

const API_URL = 'http://localhost:8888/v1/api';

let API_URL_Order = API_URL + '/orders';

const ApiOrder = {
  QueryOrderList(patient_id) {
    return axios.get(API_URL_Order + `?patient_id=${patient_id}`)
      .then(res => (res.data.payload))
      .catch(err => {
        throw new Error(`Error in ApiOrder.QueryOrderList: ${err.message}: ${err.response.data.msg}`);
      });
  },

  CreateOrder(dto) {
    return axios.post(API_URL_Order, dto, {headers: {'Content-Type': 'application/json'}})
      .then(res => (res.data.payload.id))
      .catch(err => {
        throw new Error(`Error in ApiOrder.CreateOrder: ${err.message}: ${err.response.data.msg}`);
      });
  },

  UpdateOrderInfo(dto) {
    return axios.patch(API_URL_Order + `/${dto.id}`, dto, {headers: {'Content-Type': 'application/json'}})
      .catch(err => {
        throw new Error(`Error in ApiOrder.UpdateOrder: ${err.message}: ${err.response.data.msg}`);
      });
  },
};

export default ApiOrder;
