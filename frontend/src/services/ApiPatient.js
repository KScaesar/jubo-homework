import axios from 'axios';

const API_URL_Patient = process.env.NEXT_PUBLIC_API_URL + '/patients';

const ApiPatient = {
  QueryPatientList() {
    return axios.get(API_URL_Patient)
      .then(res => (res.data.payload))
      .catch(err => {
        throw new Error(`Error in ApiPatient.QueryPatientList: ${err.message}: ${err.response.data.msg}`);
      });
  },
};

export default ApiPatient;
