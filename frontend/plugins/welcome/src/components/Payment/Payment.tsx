import React, { useState, useEffect} from 'react';
import { Link as RouterLink } from 'react-router-dom';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Alert } from '@material-ui/lab';
import Swal from 'sweetalert2'; // alert
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import { deepOrange } from '@material-ui/core/colors';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  Link,
} from '@material-ui/core';
import { makeStyles} from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import { EntInsurance } from '../../api/models/EntInsurance';
import { EntBank } from '../../api/models/EntBank';
import { EntMember } from '../../api/models/EntMember';
import { EntMoneytransfer } from '../../api/models/EntMoneytransfer';




// css style 
const useStyles = makeStyles(theme => ({
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
  root: {
    display: 'flex',
    '& > *': {
      margin: theme.spacing(1),
    },
  },
  orange: {
    color: theme.palette.getContrastText(deepOrange[500]),
    backgroundColor: deepOrange[500],
  },
  secondary: {
    main: '#f44336',
  },
}));
 
const HeaderCustom = {
  minHeight: '50px',
};

export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [errormessege, setErrorMessege] = useState(String);
  const [alerttype, setAlertType] = useState(String);

  const [errorAccountNumber, setErrorAccountNumber] = useState(true);
  const [errorPhoneNumber, setErrorPhoneNumber] = useState(true);
  const [errorPrice, setErrorPrice] = useState(true);


  const [insurances, setInsurances] = useState<EntInsurance[]>([]);
  const [banks, setBanks] = useState<EntBank[]>([]);
  const [members, setMembers] = useState<EntMember[]>([]);
  const [moneytransfers, setMoneytransfers] = useState<EntMoneytransfer[]>([]);
  const [status, setStatus] = useState(false);
  const [loading, setLoading] = useState(true);
 
 
  const [insuranceid, setInsuranceid] = useState(Number);
  const [bankid, setBankid] = useState(Number);
  const [memberid, setMemberid] = useState(Number);
  const [moneytransferid, setMoneytransferid] = useState(Number);
  const [accountname, setaccountName] = useState(String);
  const [accountnumber, setaccountNumber] = useState(String);
  const [phonenumber, setphoneNumber] = useState(String);
  const [price, setPrice] = useState(Number);
  const [transfertime, settransferTime] = useState(String);
 
  useEffect(() => {
 
    const getInsurances = async () => {
 
      const r = await api.listInsurance({ limit: 10, offset: 0 });
      setLoading(false);
      setInsurances(r);
    };
    getInsurances();
 
    const getBanks = async () => {
 
    const b = await api.listBank({ limit: 10, offset: 0 });
      setLoading(false);
      setBanks(b);
    };
    getBanks();
 
    const getMembers = async () => {
 
     const j = await api.listMember({ limit: 10, offset: 0 });
       setLoading(false);
       setMembers(j);
     };
     getMembers();
 

    const getMoneytransfers = async () => {
 
    const mn = await api.listMoneytransfer({ limit: 10, offset: 0 });
       setLoading(false);
       setMoneytransfers(mn);
     };
     getMoneytransfers();

     const data = localStorage.getItem("memberdata");
     if (data) {
     setMemberid(Number(localStorage.getItem("memberdata")));
     }

  }, [loading]);
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

 
  const createPayment = async () => {
    if ((accountname != "") && (accountnumber != "") && (bankid != null) && (insuranceid != null) 
    && (memberid != null) && (moneytransferid != null) && (transfertime != "") ) {
      const payment = {
          accountName    : accountname,
          accountNumber  : accountnumber,
          bankID         : bankid,
          insuranceID    : insuranceid,
          memberID       : memberid,
          moneytransferID : moneytransferid,
          phoneNumber    :  phonenumber,
          price          :  Number(price),
          transferTime   : transfertime + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
      };
      console.log(payment);
        const apiUrl = 'http://localhost:8080/api/v1/payments';
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payment),
        };
        fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if (data.status === true) {
                    setErrorMessege("บันทึกข้อมูลสำเร็จ");
                    setAlertType("success");

                }
                else {
                    ErrorCaseCheck(data.error.Name);
                    setAlertType("error");
                }
            });
        }
        else {
            ErrorCaseCheck("เวลาไม่ได้ใส่");
            setAlertType("error");
        }

        setStatus(true);
    };

    const ErrorCaseCheck = (casename: string) => {
      ValidateaccountNumber(accountnumber);
      Validateprice(Number(price));
      ValidatephoneNumber(phonenumber);
      if (casename == "account_number") { setErrorMessege("หมายเลขบัญชีไม่ครบ 10 หลัก"); }
      else if (casename == "price") { setErrorMessege("รูปแบบจำนวนเบี้ยประกันไม่ถูกต้อง"); }
      else if (casename == "phone_number") { setErrorMessege("เบอร์ติดต่อไม่ครบ 10 หลัก"); }
      else { setErrorMessege("บันทึกไม่สำเร็จ"); }
  }

 
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
     ValidateaccountNumber(event.target.value as string)
    };

    const phone_number_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setphoneNumber(event.target.value as string);
      ValidatephoneNumber(event.target.value as string)
     };

     const price_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setPrice(event.target.value as number);
      Validateprice(event.target.value as number)
     };
    
   
    const ValidateaccountNumber = (value : string) => {
      return value.length == 10 ?  setErrorAccountNumber(true) : setErrorAccountNumber(false);
  }
  
     const ValidatephoneNumber = (value : string) => {
    return value.length == 10 ? setErrorPhoneNumber(true) : setErrorPhoneNumber(false);
  }

  
  const Validateprice = (value : number) => {
    value > 0 ? setErrorPrice(true) : setErrorPrice(false);
  } 

  return (
    <Page theme={pageTheme.library}>
      <Header style={HeaderCustom} title={`ระบบชำระเบี้ยประกัน`}>
      </Header>
      <Content>
      <ContentHeader title="กรุณากรอกข้อมูลการชำระเงิน">
      
      </ContentHeader>
      {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>สมาชิกระบบประกันสุขภาพ</div>
            </Grid>
            <Grid item xs={8}>
            <TextField id="outlined-basic" 
              style={{ width: 300}}
              name = "member"
              variant="outlined"
              value={members.filter((filter: EntMember) => filter.id == memberid).map((item: EntMember) => `${item.memberEmail}`)}
              className={classes.textField}
              InputLabelProps={{
                shrink: true,
              }}
              
              />
            </Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>ชื่อ-สกุล</div>
            </Grid>
            <Grid item xs={8}>
            <TextField id="outlined-basic" 
              style={{ width: 300}}
              name = "account_name"
              variant="outlined"
              className={classes.textField}
              InputLabelProps={{
                shrink: true,
              }}
              onChange={account_name_handleChange}
              />
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>หมายเลขกรมธรรม์</div>
            </Grid>
            <Grid item xs={8}>
            <FormControl variant="outlined" className={classes.formControl}>
                
                <InputLabel>เลือกหมายเลขกรมธรรม์</InputLabel>
                <Select
                  name="insurance"
                  value={insuranceid || ''}
                  onChange={insurance_id_handleChange}
                >
                  {insurances.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.id}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ประเภทเบี้ยประกันภัย</div>
            </Grid>
            <Grid item xs={8}>
            <TextField id="outlined-basic" 
              style={{ width: 300}}
              name = "insuranceproduct"
              variant="outlined"
              value={insurances.filter((filter: EntInsurance) => filter.id == insuranceid).map((item: EntInsurance) => `${item.edges?.product?.productName}`)}
              className={classes.textField}
              InputLabelProps={{
                shrink: true,
              }}
              
              />
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ช่องทางการชำระเงิน</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกช่องทางการชำระเงิน</InputLabel>
                <Select
                  name="moneytransfer"
                  value={moneytransferid|| ''}
                  onChange={moneytransfer_id_handleChange}
                >
                  {moneytransfers.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.moneytransferType}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ธนาคาร</div>
            </Grid>

            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกธนาคาร</InputLabel>
                <Select
                  name="bank"
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
              <div className={classes.paper}>ป้อนเลขที่บัญชี</div>
            </Grid>
            <Grid item xs={8}>
            <TextField
              id="outlined-basic" 
              type="string"
              variant="outlined"
              style={{ width: 300 }}
              value={accountnumber}
              helperText={errorAccountNumber? "" : "กรอกเลขที่บัญชีให้ครบถ้วน"}
              error={errorAccountNumber? false : true}
              onChange={account_number_handleChange}
              />
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>จำนวนเบี้ยประกัน</div>
            </Grid>
            <Grid item xs={8}>
            <TextField
              id="price"
              name = "price"
              variant="outlined"
              value={price}
              helperText={errorPrice? "" : "กรอกรูปแบบจำนวนเงินให้ถูกต้อง"}
              error={errorPrice? false : true}
              onChange={price_handleChange}
              style={{ width: 300 }}
              />
            </Grid>
            
            <Grid item xs={4}>
              <div className={classes.paper}>เบอร์ติดต่อ</div>
            </Grid>
            <Grid item xs={8}>
            <TextField
              id="phonenumber"
              name = "phonenumber"
              variant="outlined"
              value={phonenumber}
              helperText={errorPhoneNumber? "" : "กรุณากรอกหมายเลขโทรศัพท์ให้ครบถ้วน"}
              error={errorPhoneNumber? false : true}
              onChange={phone_number_handleChange}
              style={{ width: 300 }}
              />
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>เวลา</div>
            </Grid>
            <Grid item xs={8}>
            <form className={classes.container} noValidate>
                <TextField
                  label="เลือกเวลา"
                  name="transfertime"
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