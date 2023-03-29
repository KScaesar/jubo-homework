import React, {useEffect, useState} from 'react';
import {Box, Dialog, DialogActions, DialogContent, DialogTitle, Grid, IconButton, TextField, Typography} from '@mui/material';
import AddBoxIcon from '@mui/icons-material/AddBox';
import DisabledByDefaultOutlinedIcon from '@mui/icons-material/DisabledByDefaultOutlined';
import EditRoundedIcon from '@mui/icons-material/EditRounded';
import CheckCircleRoundedIcon from '@mui/icons-material/CheckCircleRounded';
import ApiOrder from "@/services/ApiOrder";

export default function OrderDialog({open, onClose, patient}) {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    const go = async () => {
      const view = {mode: 'r', click: 0};
      let {list} = await ApiOrder.QueryOrderList(patient.id);
      setOrders(list.map((order) => ({...order, ...view})))
    }
    go()
  }, [patient.id,])

  const handleApiSaveOrder = (buttonEvent, idx) => {
    if (buttonEvent && buttonEvent.type === '"click"') {
      buttonEvent.preventDefault()
    }

    let target = orders.find((order, i) => (i === idx));
    if (target.mode === 'r') {
      return
    }

    const mode = target.mode
    target.mode = 'r'
    const go = async () => {
      if (mode === 'u') {
        await ApiOrder.UpdateOrderInfo(target)
      } else if (mode === 'c') {
        let id = await ApiOrder.CreateOrder(target)
        target.id = id
      }
      setOrders([...orders])
    }
    go()
  }

  const handleAddViewOrder = (patient) => {
    let newOrder = {
      id: new Date().getTime().toString(),
      message: 'edit message...',
      patient_id: patient.id,

      mode: 'c',
      firstClick: true,
    };
    newOrder.mode = 'c'
    setOrders([newOrder, ...orders]);
  };

  const handleChangeSymbol = (event, idx) => {
    event.preventDefault()
    let target = orders.find((order, i) => (i === idx));
    target.mode = 'u'
    setOrders([...orders])
  }

  function handleClickText(idx) {
    let target = orders.find((order, i) => (i === idx));
    if (target.mode === 'c' && target.firstClick) {
      target.message = ''
      target.firstClick = false
    }
    setOrders([...orders])
  }

  function handleEnter(event, idx) {
    if (event.key === "Enter") {
      handleApiSaveOrder(null, idx)
    }
  }

  function handleTyping(event, idx) {
    let target = orders.find((order, i) => (i === idx))
    target.message = event.target.value
    setOrders([...orders])
  }

  return (
    <Dialog open={open} onClose={onClose} fullWidth={true} sx={{height: '100%'}}>
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
      <DialogContent>
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
                onMouseDown={() => handleClickText(idx)}
                onKeyDown={(event) => handleEnter(event, idx)}
              />
            </Grid>
            <Grid item xs={1}>
              <MyButton
                mode={order.mode}
                idx={idx}
                onClick={order.mode === 'r' ? handleChangeSymbol : handleApiSaveOrder}/>
            </Grid>
          </Grid>
        ))}
      </DialogContent>
    </Dialog>
  );
};

const MyButton = (props) => {
  return props.mode === 'r' ? (
    <IconButton color="primary" {...props} onClick={(event) => props.onClick(event, props.idx)}>
      <Box id={props.id} sx={{fontSize: 30}}>
        <EditRoundedIcon id={props.id}/>
      </Box>
    </IconButton>
  ) : (
    <IconButton color="secondary" {...props} onClick={(event) => props.onClick(event, props.idx)}>
      <Box component={CheckCircleRoundedIcon} sx={{fontSize: 30}}/>
    </IconButton>
  );
};