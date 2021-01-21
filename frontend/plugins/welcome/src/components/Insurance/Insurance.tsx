import React, { useState, useEffect} from 'react';
import { Link as RouterLink } from 'react-router-dom';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Alert } from '@material-ui/lab';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
  Link,
} from '@material-ui/core';
import { makeStyles} from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import { EntInsurance } from '../../api/models/EntInsurance';
import { EntBank } from '../../api/models/EntBank';
import { EntMember } from '../../api/models/EntMember';
import { EntMoneyTransfer } from '../../api/models/EntMoneyTransfer';


// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(1),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
  datetimepicker:{
    width: 300,
  },
}));
 
const HeaderCustom = {
  minHeight: '50px',
};

export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();


  const [insurances, setInsurances] = useState<EntInsurance[]>([]);
  const [banks, setBanks] = useState<EntBank[]>([]);
  const [members, setMembers] = useState<EntMember[]>([]);
  const [moneytransfers, setMoneytransfers] = useState<EntMoneyTransfer[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
 
 
  const [insuranceid, setInsuranceid] = useState(Number);
  const [bankid, setBankid] = useState(Number);
  const [memberid, setMemberid] = useState(Number);
  const [moneytransferid, setMoneytransferid] = useState(Number);
  const [accountname, setaccountName] = useState(String);
  const [accountnumber, setaccountNumber] = useState(String);
  const [transfertime, settransferTime] = useState(String);
 
  useEffect(() => {
 
    const getInsurances = async () => {
 
      const r = await api.listInsurance({ limit: 10, offset: 0 });
      setLoading(false);
      setInsurances(r);
    };
    getInsurances();
 
    const getBanks = async () => {
 
    const u = await api.listBank({ limit: 10, offset: 0 });
      setLoading(false);
      setBanks(u);
    };
    getBanks();
 
    const getMembers = async () => {
 
     const j = await api.listMember({ limit: 10, offset: 0 });
       setLoading(false);
       setMembers(j);
     };
     getMembers();

    const getMoneytransfers = async () => {
 
    const j = await api.listMoneyTransfer({ limit: 10, offset: 0 });
       setLoading(false);
       setMoneytransfers(j);
     };
     getMoneytransfers();

  }, [loading]);
 
  const createPayment = async () => {
      const payment = {
        transferTime   : transfertime + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
        bankID         : bankid,
        memberID       : memberid,
        insuranceID    : insuranceid,
        moneytarnferID : moneytransferid,
        acoountName    : accountname,
        accountNumber  : accountnumber,
      }
      console.log(payment);

    const res:any = await api.createPayment({ payment : payment});
    setStatus(true);
    if (res.id != ''){
      setAlert(true);
    } else {
      setAlert(false);
    }
 
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };
 
   const bank_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setBankid(event.target.value as number);
    };

   const insurance_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setInsuranceid(event.target.value as number);
    };

   const member_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setMemberid(event.target.value as number);
    };

   const moneytransfer_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setMoneytransferid(event.target.value as number);
    };

   const transfer_time_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     settransferTime(event.target.value as string);
    };

   const account_name_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setaccountName(event.target.value as string);
    };

   const account_number_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setaccountNumber(event.target.value as string);
    };

 
  return (
    <Page theme={pageTheme.library}>
      <Header style={HeaderCustom} title={`ระบบบริการรถพยาบาล`}>
        <Avatar alt="Teerasak Supavaha" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10 }}>Teerasak Supawaha</div>
        <td></td>
        <Link component={RouterLink} to="/">logout
         </Link>
      </Header>
      <Content>
      <ContentHeader title="ใบขอใช้บริการรถพยาบาล">
      <Link component={RouterLink} to="/results">
           <Button variant="contained" color="secondary">
             ผลการบันทึกข้อมูล
           </Button>
         </Link>
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 This is a success alert — check it out!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
         ) : null}
      </ContentHeader>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>ชื่อ-สกุล (ผู้ร้องขอ)</div>
            </Grid>
            <Grid item xs={8}>
            <FormControl variant="outlined" className={classes.formControl}>
            <TextField id="outlined-basic" 
              style={{ width: 300}}
              label ="ป้อนชื่อ-สกุล"
              name = "requestorname"
              variant="outlined"
              value={accountnumber || ''}// (undefined|| ") = "
              className={classes.textField}
              InputLabelProps={{
                shrink: true,
              }}
              onChange={account_number_handleChange} />
              </FormControl>
            </Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>หมายเลขรถพยาบาล</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                
                <InputLabel>เลือกหมายเลขรถพยาบาล</InputLabel>
                <Select
                  name="car"
                  value={bankid || ''}
                  onChange={bank_id_handleChange}
                >
                  {banks.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.bankType}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ลักษณะงาน</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกลักษณะงาน</InputLabel>
                <Select
                  name="job"
                  value={memberid || ''}
                  onChange={member_id_handleChange}
                >
                  {members.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.memberName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>เจ้าหน้าที่จัดการรถพยาบาล</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกเจ้าหน้าที่จัดการรถพยาบาล</InputLabel>
                <Select
                  name="user"
                  value={insuranceid || ''}
                  onChange={insurance_id_handleChange}
                >
                  {insurances.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.edges?.product?.productName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>สถานที่</div>
            </Grid>
            <Grid item xs={8}>
            <TextField id="outlined-basic" 
              style={{ width: 300}}
              name = "lacation"
              label="ป้อนสถานที่" 
              variant="outlined"
              value={location}
              className={classes.textField}
              InputLabelProps={{
                shrink: true,
              }}
              onChange={account_name_handleChange}
              />
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>เวลา</div>
            </Grid>
            <Grid item xs={8}>
            <form className={classes.container} noValidate>
                <TextField
                  label="เลือกเวลา"
                  name="added"
                  type="datetime-local"
                  value={transfertime || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={transfer_time_handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={4}></Grid>
            <Grid item xs={8}>
              <Button
                variant="contained"
                color="secondary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={() => {
                  createPayment();
                }}
              >
                บันทึกข้อมูล
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};