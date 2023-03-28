import axios from 'axios';

const API_URL = 'http://localhost:8888/v1/api';

let API_URL_Order = API_URL + '/orders';

const ApiOrder = {
  QueryOrderList(patient_id) {
    return axios.get(API_URL_Order + `?patient_id=${patient_id}`)
      .then(res => (res.data.payload))
      .catch(err => {
        console.log(err)
        throw new Error(`Error in ApiOrder.QueryOrderList: ${err.message}`);
      });
  },

  CreateOrder(dto) {
    return axios.post(API_URL_Order, dto, {headers: {'Content-Type': 'application/json'}})
      .catch(err => {
        console.log(err)
        throw new Error(`Error in ApiOrder.CreateOrder: ${err.message}`);
      });
  },

  UpdateOrderInfo({id, dto}) {
    return axios.patch(API_URL_Order + `/${id}`, dto, {headers: {'Content-Type': 'application/json'}})
      .catch(err => {
        console.log(err)
        throw new Error(`Error in ApiOrder.UpdateOrder: ${err.message}`);
      });
  },
};

export default ApiOrder;
