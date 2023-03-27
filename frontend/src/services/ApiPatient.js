import axios from 'axios';

const API_URL = 'http://localhost:8888/v1/api';

let API_URL_Patient = API_URL + '/patients';

const ApiPatient = {
  QueryPatientList() {
    return axios.get(API_URL_Patient)
      .then(res => (res.data.payload))
      .catch(err => {
        console.log(err)
        throw new Error(`Error in ApiPatient.QueryList: ${err.message}`);
      });
  },
};

export default ApiPatient;
