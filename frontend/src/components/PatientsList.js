import React, {useEffect, useState} from 'react';
import {List, ListItemButton, ListItemIcon, ListItemText} from '@mui/material';
import PermIdentityIcon from '@mui/icons-material/PermIdentity';
import ApiPatient from "@/services/ApiPatient";
import OrderDialog from "@/components/OrderDialog";

const PatientsList = () => {
  const [selectedPatient, setSelectedPatient] = useState(null);
  const [dialogOpen, setDialogOpen] = useState(false);
  const [patients, setPatients] = useState([]);

  useEffect(() => {
    async function go() {
      try {
        const {list} = await ApiPatient.QueryPatientList();
        setPatients(list);
      } catch (error) {
        console.error(error);
      }
    }

    go();
  }, [])

  const handlePatientClick = (patient) => {
    setSelectedPatient({...patient});
    setDialogOpen(true);
  };

  const handleDialogClose = () => {
    setSelectedPatient(null);
    setDialogOpen(false);
  };

  return (
    <div>
      <List sx={{backgroundColor: '#f0f0f0'}}>
        {patients.map((patient) => (
          <ListItemButton key={patient.id} onClick={() => handlePatientClick(patient)}>
            <ListItemIcon>
              <PermIdentityIcon/>
            </ListItemIcon>
            <ListItemText primary={patient.name}/>
          </ListItemButton>
        ))}
      </List>
      {selectedPatient ?
        <OrderDialog
          open={dialogOpen}
          onClose={handleDialogClose}
          patient={selectedPatient}
        /> :
        null
      }
    </div>
  );
};

export default PatientsList;
