import React, {useEffect, useState} from 'react';
import {Box, Dialog, DialogActions, DialogContent, DialogTitle, Grid, IconButton, TextField, Typography} from '@mui/material';
import AddBoxIcon from '@mui/icons-material/AddBox';
import DisabledByDefaultOutlinedIcon from '@mui/icons-material/DisabledByDefaultOutlined';
import EditRoundedIcon from '@mui/icons-material/EditRounded';
import CheckCircleRoundedIcon from '@mui/icons-material/CheckCircleRounded';
import ApiOrder from "@/services/ApiOrder";

export default function OrderDialog({open, onClose, patient}) {
  const [orders, setOrders] = useState([]);
  const [selectedOrderIdx, setSelectedOrderIdx] = useState(-1)

  useEffect(() => {
    const go = async () => {
      const view = {mode: 'r', click: 0};
      let {list} = await ApiOrder.QueryOrderList(patient.id);
      setOrders(list.map((order) => ({...order, ...view})))
    }
    go()
  }, [patient.id])

  const handleAddViewOrder = (patient) => {
    let newOrder = {
      id: new Date().getTime().toString(),
      message: 'edit message...',
      patient_id: patient.id,

      mode: 'c',
      click: 0,
    };
    newOrder.mode = 'c'
    setOrders([newOrder, ...orders]);
  };

  const handleConfirmOrder = (id, msg) => {

  }

  function handleClickText(idx) {
    let target = orders.find((order, i) => (i === idx));
    if (target.mode === 'c' && target.click === 0) {
      target.message = ''
      target.click++
    }
    setOrders([...orders])
  }

  function handleTyping(event, idx) {
    let target = orders.find((order, i) => (i === idx))
    target.message = event.target.value
    setOrders([...orders])
  }

  return (
    <Dialog open={open} onClose={onClose} fullWidth={true}>
      <DialogTitle style={{display: 'flex', alignItems: 'center', justifyContent: 'space-between',}}>
        <Typography variant="p">{patient ? patient.name : ''}</Typography>
        <DialogActions>
          <IconButton color="primary" onClick={() => handleAddViewOrder(patient)}>
            <Box component={AddBoxIcon} sx={{fontSize: 30}}/>
          </IconButton>
          <IconButton color="secondary" onClick={onClose}>
            <Box component={DisabledByDefaultOutlinedIcon} sx={{fontSize: 30}}/>
          </IconButton>
        </DialogActions>
      </DialogTitle>
      <DialogContent sx={{minHeight: '1000px'}}>
        {orders.map((order, idx) => (
          <Grid key={order.id} container spacing={2} alignItems="center" justifyContent="center">
            <Grid item xs={11}>
              <TextField
                key={order.id}
                value={order.message}
                margin="normal"
                fullWidth
                multiline={true}
                rows={2}
                disabled={order.mode === 'r'}
                onChange={(event) => handleTyping(event, idx)}
                onClick={() => handleClickText(idx)}
              />
            </Grid>
            <Grid item xs={1}>
              <MyButton mode={order.mode}/>
            </Grid>
          </Grid>
        ))}
      </DialogContent>
    </Dialog>
  );
};

const MyButton = (props) => {
  return props.mode === 'r' ? (
    <IconButton color="primary"  {...props}>
      <Box component={EditRoundedIcon} sx={{fontSize: 30}}/>
    </IconButton>
  ) : (
    <IconButton color="secondary"  {...props}>
      <Box component={CheckCircleRoundedIcon} sx={{fontSize: 30}}/>
    </IconButton>
  )
};