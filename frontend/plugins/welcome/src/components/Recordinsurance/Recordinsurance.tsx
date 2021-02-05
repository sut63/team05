import React, { FC, useEffect,useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert


import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
} from '@material-ui/core'
import { DefaultApi } from '../../api/apis';
import { EntProduct } from '../../api/models/EntProduct';
import { EntMember } from '../../api/models/EntMember';
import { EntHospital } from '../../api/models/EntHospital';
import { EntOfficer } from '../../api/models/EntOfficer';
import { EntAmountpaid } from '../../api/models/EntAmountpaid';
import { Alert } from '@material-ui/lab';
import { ContentHeader } from '@backstage/core';
import InputAdornment from '@material-ui/core/InputAdornment';
import OutlinedInput from '@material-ui/core/OutlinedInput';
import Autocomplete from '@material-ui/lab/Autocomplete';
import { Link as RouterLink } from 'react-router-dom';
import { Link } from '@material-ui/core';
import { Theme, createStyles } from '@material-ui/core/styles';


// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
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
}));


export default function Create() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [alerttype, setAlertType] = useState(String);
   
    const [members, setMembers] = useState<EntMember[]>([]);
    const [products, setProducts] = useState<EntProduct[]>([]);
    const [amountpaids, setAmountpaids] = useState<EntAmountpaid[]>([]);
    const [officers, setOfficers] = useState<EntOfficer[]>([]);
    const [hospitals, setHospitals] = useState<EntHospital[]>([]);
  
    const [status, setStatus] = useState(false);
    const [alert, setAlert] = useState(true);
    const [loading, setLoading] = useState(true);
    
    const [officerID, setOfficerID] = React.useState(Number);
    const [memberid, setMemberid] = React.useState(Number);
    const [productid, setProductid] = React.useState(Number);
    const [amountpaidid, setAmountpaidid] = React.useState(Number);
    const [officerid, setOfficerid] = React.useState(Number);
    const [hospitalid, setHospitalid] = useState(Number);
 
    //const [recordinsurance_age, setRecordinsuranceAge] = useState(Number);
    const [number_of_days_of_treat, setNumberOfDaysOfTreat] = useState(Number);
    const [recordinsurance_contact, setRecordinsuranceContact] = useState(String);
    //const [recordinsurance_time, setRecordinsuranceTime] = useState(String);
    const [recordinsurance_address, setRecordinsuranceAddress] = useState(String);

    const [number_of_days_of_treaterror, setNumberOfDaysOfTreatError] = React.useState('');
    const [recordinsurance_contacterror, setContactError] = React.useState('');
    const [recordinsurance_addresserror, setAddressError] = React.useState('');


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
    
    const setErrorMessege = (icon: any , title: any) => {
      Toast.fire({
        icon: icon,
        title: title,
      })
    }
    
  useEffect(() => {
 
    const getProducts = async () => {
 
      const cn = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(cn);
    };
    getProducts();
 
    const getMembers = async () => {
 
    const u = await api.listMember({ limit: 10, offset: 0 });
      setLoading(false);
      setMembers(u);
    };
    getMembers();
 
    const getHospitals = async () => {
 
     const s = await api.listHospital({ limit: 10, offset: 0 });
       setLoading(false);
       setHospitals(s);
     };
     getHospitals();

     const getOfficers = async () => {
 
      const st = await api.listOfficer({ limit: 10, offset: 0 });
        setLoading(false);
        setOfficers(st);
      };
      getOfficers();

      const getAmountpaids = async () => {
 
        const am = await api.listAmountpaid({ limit: 10, offset: 0 });
          setLoading(false);
          setAmountpaids(am);
        };
        getAmountpaids();

    const dataa = localStorage.getItem("officerdata");
    if (dataa) {
    setOfficerID(Number(localStorage.getItem("officerdata")));
    setLoading(false);
    }

    const checkJobPosition = async () => {
      const jobdata = JSON.parse(String(localStorage.getItem("positiondata")));
      setLoading(false);
      if (jobdata != "พนักงานบริษัทประกันสุขภาพ" ) {
        localStorage.setItem("officerdata",JSON.stringify(null));
        localStorage.setItem("positiondata",JSON.stringify(null));
        history.pushState("","","./Officerlongin");  
        window.location.reload(false);       
      }
      else{
          setOfficerID(Number(localStorage.getItem("officerdata")))
      }
    }
  checkJobPosition();
 
  }, [loading]);
     


   const product_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
   setProductid(event.target.value as number);
    };

   const member_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setMemberid(event.target.value as number);
    };

    const hospital_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setHospitalid(event.target.value as number);
     };

     const officer_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setOfficerid(event.target.value as number);
     };

     const amountpaid_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setAmountpaidid(event.target.value as number);
     };

     /*const recordinsuranceAge_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setRecordinsuranceAge(event.target.value as number);
     };*/

     const numberOfDaysOfTreat_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
      const { value } = event.target;
        const validateValue = value
        checkpattern('number_of_days_of_treat', validateValue)
        setNumberOfDaysOfTreat(event.target.value as number);
    };

     const recordinsuranceContact_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
      const { value } = event.target;
        const validateValue = value
        checkpattern('recordinsurance_contact', validateValue)
        setRecordinsuranceContact(event.target.value as string);
    };

     /*const recordinsuranceTime_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setRecordinsuranceTime(event.target.value as string);
     };*/

     const recordinsuranceAddress_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
      const { value } = event.target;
        const validateValue = value
        checkpattern('recordinsurance_address', validateValue)
        setRecordinsuranceAddress(event.target.value as string);
    };


    const validateAddress = (val: string) => {
      return val.match("^[ก-๙0-9a-zA-Z- ./\\s]+$");
   }
   
   const validateContact = (val: string) => {
    return val.match("^[0]\\d")
  }

   const validateNumberOfDaysOfTreat = (val: number) => {
    return val <= 30 && val >= 0 ? true : false 
}

   const checkpattern = (id: string, value:string) => {
    console.log(value);
    switch(id) {
      case 'recordinsurance_contact' :
        validateContact(value) ? setContactError('') : setContactError('กรุณากรอกเบอร์โทรติดต่อให้ถูกต้อง');
        return;
      case 'recordinsurance_address':
        validateAddress(value) ? setAddressError('') : setAddressError('กรุณากรอกที่อยู่ให้ถูกต้อง ไม่สามารถเป็นค่าว่างได้');
        return;
        case 'number_of_days_of_treat':
          validateNumberOfDaysOfTreat(Number(value)) ? setNumberOfDaysOfTreatError('') : setNumberOfDaysOfTreatError('กรุณากรอกจำนวนวันให้ถูกต้อง');
      return;
      default:
        return;
    }
  }
     

  const CreateRecordinsurance = async () => {
    /*if ((amountpaidid != null) && (hospitalid != null) && (memberid != null) && (officerid != null) && (productid != null) && (number_of_days_of_treat != 0) && (number_of_days_of_treat != null )
    && (recordinsurance_address != null) && (recordinsurance_address != "") && (recordinsurance_contact != null) && (recordinsurance_contact != "")){*/
      
    if ((amountpaidid != null) && (hospitalid != null) && (memberid != null) && (officerid != null) && (productid != null)){

      const apiUrl = 'http://localhost:8080/api/v1/recordinsurances';
      const recordinsurance = {

         amountpaidID      : amountpaidid,
         hospitalID        : hospitalid,
         memberID          : memberid,
         numberOfDaysOfTreat  : Number(number_of_days_of_treat),
         officerID         : officerID,
         productID         : productid,
         recordinsuranceAddress : recordinsurance_address,
         recordinsuranceContact : recordinsurance_contact,
         //recordinsuranceTime    : recordinsurance_time + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
      }
      console.log(recordinsurance);
    
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(recordinsurance),
      };
          fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
              console.log(data);
              if (data.status === true) {
                Toast.fire({
                  icon: 'success',
                  title: 'บันทึกข้อมูลสำเร็จ',
                });
              }
              else {
                ErrorCaseCheck(data.error.Name);
                setAlertType("error");
              }  
            });
        }
        else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
          setAlertType("error");
        }
          setStatus(true);
      
      };

      const ErrorCaseCheck = (field: string) => {
        switch(field) {
          case 'recordinsurance_contact':
            setErrorMessege("error","เบอโทรไม่ถูกต้องต้องขึ้นต้นด้วย 0 และครบ 10 ตัว");
            return;
          case 'number_of_days_of_treat':
            setErrorMessege("error","จำนวนวันที่ระบุไม่เกิน30วัน");
            return;
          case 'recordinsurance_address':
            setErrorMessege("error","ระบุที่อยู่ที่ติดต่อได้");
            return;
          default:
            setErrorMessege("error","บันทึกข้อมูลไม่สำเร็จ");
            return;
        }
      }


    
    

    return (
      <Page theme={pageTheme.tool}>
        <Header style={HeaderCustom} title={`ระบบบันทึกข้อมูลสิทธิประกันสุขภาพ`}>
        </Header>
        <Content>
        <ContentHeader title="ทำการบันทึกข้อมูลสิทธิประกันสุขภาพ" >
        <Button
                    style={{ marginLeft: 10 }}
                    component={RouterLink}
                    to="/recordinsurancesearch"
                    variant="contained"
                    color="secondary"
                  >
                    ค้นหาการบันทึก
                </Button>
        </ContentHeader>

          <Container maxWidth="sm">
          <Grid container spacing={3}>
          <Grid item xs={10}></Grid>
          <Grid item xs={3}>
                <div className={classes.paper}>สมาชิกระบบประกันสุขภาพ</div>
              </Grid>
              <Grid item xs={8}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>เลือกสมาชิกระบบประกันสุขภาพ</InputLabel>
                  <Select
                    name="member"
                    value={memberid || ''} // (undefined || '') = ''
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


              
              <Grid item xs={3}>
                <div className={classes.paper}>ผลิตภัณฑ์</div>
              </Grid>
              <Grid item xs={8}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>เลือกผลิตภัณฑ์</InputLabel>
                  <Select
                    name="product"
                    value={productid || ''} // (undefined || '') = ''
                    onChange={product_id_handleChange}
                  >
                    {products.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.productName}
                        </MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={3}>
              <div className={classes.paper}>เงินที่ประกันภัยจ่าย</div>
            </Grid>
            <Grid item xs={8}>
            <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>฿</InputLabel>
                <Select
                    name="amountpaid"
                    value={amountpaidid || ''} // (undefined || '') = ''
                    onChange={amountpaid_id_handleChange}
                  >
                    {amountpaids.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.amountpaidMoney}
                        </MenuItem>
                      );
                    })}
                  </Select>
              </FormControl>
            </Grid>
  
              <Grid item xs={3}>
                <div className={classes.paper}>โรงพยาบาล</div>
              </Grid>
              <Grid item xs={8}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>เลือกโรงพยาบาล</InputLabel>
                  <Select
                    name="hospital"
                    value={hospitalid || ''} // (undefined || '') = ''
                    onChange={hospital_id_handleChange}
                  >
                    {hospitals.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.hospitalName}
                        </MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={3}>
              <div className={classes.paper}>จำนวนวันในการรักษา</div>
            </Grid>
            <Grid item xs={8}>
            <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                id="number_of_days_of_treat" 
                type='number' InputLabelProps={{
                shrink: true,
                }} 
                label="" variant="outlined"
                onChange={numberOfDaysOfTreat_handleChange}
                value={number_of_days_of_treat || ''}
                />
              </FormControl>
            </Grid>
  
              <Grid item xs={3}>
                <div className={classes.paper}>พนักงานบริษัทประกันสุขภาพ</div>
              </Grid>
              <Grid item xs={8}>
              <TextField id="outlined-basic" 
              style={{ width: 300}}
              name = "officer"
              variant="outlined"
              value={officers.filter((filter: EntOfficer) => filter.id == officerID).map((item: EntOfficer) => `${item.officerName}`)}
              className={classes.textField}
              InputLabelProps={{
                shrink: true,
              }}
              
              />
              </Grid>

              <Grid item xs={3}>
              <div className={classes.paper}>ที่อยู่ของลูกค้า</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  error = {recordinsurance_addresserror ? true : false}
                  id="outlined-number"
                  name="recordinsurance_address"
                  multiline
                  rows={3}
                  value={recordinsurance_address}
                  type="string"
                  onChange={recordinsuranceAddress_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                  helperText ={recordinsurance_addresserror}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เบอร์โทรติดต่อ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  error = {recordinsurance_contacterror ? true : false}
                  id="outlined-number"
                  name="recordinsurance_contact"
                  multiline
                  rows={1}
                  value={recordinsurance_contact}
                  type="string"
                  onChange={recordinsuranceContact_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                  helperText ={recordinsurance_contacterror}
                />
              </FormControl>
            </Grid>

            

              <Grid item xs={3}></Grid>
              <Grid item xs={7}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={() => {
                  CreateRecordinsurance();
                }}
              >
                บันทึกข้อมูลสิทธิประกันสุขภาพ
              </Button>
            </Grid>
          </Grid>
          </Container>
        </Content>
      </Page>
    );
  };