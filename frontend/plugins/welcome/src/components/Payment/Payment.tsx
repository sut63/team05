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
import { EntProduct } from '../../api/models/EntProduct';



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
  const [products, setProducts] = useState<EntProduct[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
 
 
  const [insuranceid, setInsuranceid] = useState(Number);
  const [bankid, setBankid] = useState(Number);
  const [memberid, setMemberid] = useState(Number);
  const [productid, setProductid] = useState(Number);
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

     const getProducts = async () => {
 
      const pr = await api.listProduct({ limit: 10, offset: 0 });
        setLoading(false);
        setProducts(pr);
      };
      getProducts();
 

    const getMoneytransfers = async () => {
 
    const mn = await api.listMoneytransfer({ limit: 10, offset: 0 });
       setLoading(false);
       setMoneytransfers(mn);
     };
     getMoneytransfers();

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
      const payment = {
        acoountName    : accountname,
        accountNumber  : accountnumber,
        bankID         : bankid,
        insuranceID    : insuranceid,
        memberID       : memberid,
        moneytarnferID : moneytransferid,
        productID      : productid,
        transferTime   : transfertime + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
      }
      console.log(payment);
      

    const res:any = await api.createPayment({ payment : payment});
    setStatus(true);
    if (res !== undefined) {
      Toast.fire({
        icon: 'success',
        title: 'บันทึกข้อมูลสำเร็จ',
      });
    } else {
      Toast.fire({
        icon: 'error',
        title: 'บันทึกข้อมูลไม่สำเร็จ',
      });
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

   const product_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setProductid(event.target.value as number);
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
      <Header style={HeaderCustom} title={`ระบบชำระเบี้ยประกัน`}>
        <Avatar alt="Teerasak Supavaha" src="../../image/account.jpg" className={classes.orange}/>
        <div style={{ marginLeft: 10 }}>Teerasak Supawaha</div>
        <td></td>
      </Header>
      <Content>
      <ContentHeader title="กรุณากรอกข้อมูลการชำระเงิน">
      <Link component={RouterLink} to="/results">
           <Button variant="contained" color="secondary">
             ผลการบันทึกข้อมูล
           </Button>
         </Link>
      </ContentHeader>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>สมาชิกระบบประกันสุขภาพ</div>
            </Grid>
            <Grid item xs={8}>
            <FormControl variant="outlined" className={classes.formControl}>
                
                <InputLabel>เลือกสมาชิกระบบประกันสุขภาพ</InputLabel>
                <Select
                  name="member"
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
              <div className={classes.paper}>เบี้ยประกัน</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                
                <InputLabel>เลือกเบี้ยประกัน</InputLabel>
                <Select

                  name="product"
                  value={productid || ''}
                  onChange={product_id_handleChange}
                >
                  {products.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.productPrice}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
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
                  value={bankid|| ''}
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
              <div className={classes.paper}>ป้อนชื่อบัญชี</div>
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
              <div className={classes.paper}>ป้อนเลขที่บัญชี</div>
            </Grid>
            <Grid item xs={8}>
            <TextField id="outlined-basic" 
              style={{ width: 300}}
              name = "account_number"
              variant="outlined"
              className={classes.textField}
              InputLabelProps={{
                shrink: true,
              }}
              onChange={account_number_handleChange}
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