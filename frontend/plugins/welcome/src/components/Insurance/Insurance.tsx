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
  Avatar,
  Button,
  FormHelperText,
} from '@material-ui/core'
import { DefaultApi } from '../../api/apis';
import Autocomplete from '@material-ui/lab/Autocomplete';
import { EntProduct } from '../../api/models/EntProduct';
import { EntMember } from '../../api/models/EntMember';
import { EntHospital } from '../../api/models/EntHospital';
import { EntOfficer } from '../../api/models/EntOfficer';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import InputAdornment from '@material-ui/core/InputAdornment';
import { Link as RouterLink } from 'react-router-dom';
import OutlinedInput from '@material-ui/core/OutlinedInput';
import { Alert } from '@material-ui/lab';
import { ContentHeader } from '@backstage/core';
import { Link } from '@material-ui/core';
import { Theme, createStyles } from '@material-ui/core/styles';

import CardMedia from '@material-ui/core/CardMedia';
import { Image1Base64Function } from '../../image/image1';

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
  media: {
    height: 0,
    marginLeft: 25,
    maxWidth: 300,
}
}));

interface Insurance {
  productid: string;
  memberid:   string;
  hospitalid:   number;
  officerid:        number;
  insurance_age:       number;
  insurance_identification:      number;
  insurance_address:      number;
  insurance_insurer: string;
  insuranceTimeFirstpays: string;
  // create_by: number;
}

export default function Insurance() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [alerttype, setAlertType] = useState(String);

  const [products, setProducts] = useState<EntProduct[]>([]);
  const [members, setMembers] = useState<EntMember[]>([]);
  const [hospitals, setHospitals] = useState<EntHospital[]>([]);
  const [officers, setOfficers] = useState<EntOfficer[]>([]);
  const [loading, setLoading] = useState(true);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 
 
  const [productid, setProductid] = useState(Number);
  const [memberid, setMemberid] = useState(Number);
  const [hospitalid, setHospitalid] = useState(Number);
  const [officerid, setOfficerid] = useState(Number);
  const [insurance_identification, setInsuranceIdentifications] = useState(String);
  const [insurance_address, setInsuranceAddress] = useState(String);
  const [insurance_insurer, setInsuranceInsurer] = useState(String);
  //const [insuranceTimeBuys, setInsuranceTimeBuy] = useState(String);
  const [insuranceTimeFirstpays, setInsuranceTimeFirstpay] = useState(String);

  const [insurance_insurererror, setInsurerError] = React.useState('');
  const [insurance_identificationerror, setIdentificationError] = React.useState('');
  const [insurance_addressserror, setAddressError] = React.useState('');

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
  
  const setErrorMessege = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }
 
  useEffect(() => {
 
    const getProducts = async () => {
 
      const pd = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(pd);
    };
    getProducts();
 
    const getMembers = async () => {
 
    const m = await api.listMember({});
      setLoading(false);
      setMembers(m);
    };
    getMembers();
 
    const getHospitals = async () => {
 
     const h = await api.listHospital({ limit: 10, offset: 0 });
       setLoading(false);
       setHospitals(h);
     };
     getHospitals();

     const getOfficers = async () => {
 
      const of = await api.listOfficer({ limit: 10, offset: 0 });
        setLoading(false);
        setOfficers(of);
      };
      getOfficers();

      const checkJobPosition = async () => {
        const jobdata = JSON.parse(String(localStorage.getItem("positiondata")));
        setLoading(false);
        if (jobdata != "สมาชิกระบบประกันสุขภาพ" ) {
          localStorage.setItem("memberdata",JSON.stringify(null));
          localStorage.setItem("positiondata",JSON.stringify(null));
          history.pushState("","","./SignIn");
          window.location.reload(false);        
        }
        else{
            setMemberid(Number(localStorage.getItem("memberdata")))
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
 
     const insuranceTimeFirstpay_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
       setInsuranceTimeFirstpay(event.target.value as string);
     };

      const insuranceIdentification_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
        const { value } = event.target;
          const validateValue = value
          checkpattern('insurance_identification', validateValue)
          setInsuranceIdentifications(event.target.value as string);
      };

      const insuranceInsurer_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
        const { value } = event.target;
          const validateValue = value
          checkpattern('insurance_insurer', validateValue)
          setInsuranceInsurer(event.target.value as string);
      };

      const insuranceAddress_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
        const { value } = event.target;
          const validateValue = value
          checkpattern('insurance_address', validateValue)
          setInsuranceAddress(event.target.value as string);
      };
      
      const validateAddress = (val: string) => {
        return val.match("^[ก-๙0-9a-zA-Z- ./\\s]+$");
     }
     
     const validateIdentification = (val: string) => {
      return val.match("^[0-9]+$");
    }

     const validateInsurer = (val: string) => {
       return val.match("^[a-zA-Z ]+$");
     }

     const checkpattern = (id: string, value:string) => {
      console.log(value);
      switch(id) {
        case 'insurance_identification' :
          validateIdentification(value) ? setIdentificationError('') : setIdentificationError('กรุณากรอกเลขประจำตัวประชาชนให้ถูกต้อง');
          return;
        case 'insurance_address':
          validateAddress(value) ? setAddressError('') : setAddressError('กรุณากรอกที่อยู่ให้ถูกต้อง');
          return;
          case 'insurance_insurer':
            validateInsurer(value) ? setInsurerError('') : setInsurerError('ระบุชื่อเป็นภาษาอังกฤษให้ถูกต้อง และตัวแรกเป็นตัวพิมพ์ใหญ่');
        return;
        default:
          return;
      }
    }


  const CreateInsurance = async () => {
    /*if ( (insurance_address != null) && (insurance_address != "") && (insurance_insurer!= null) && (insurance_insurer != "") && (insuranceTimeFirstpays != null) && (insuranceTimeFirstpays != "") && (productid != null) && (memberid != null) && (hospitalid != null) && (officerid != null) ) {*/
      if ((insuranceTimeFirstpays != null) && (insuranceTimeFirstpays != "")){
      const apiUrl = 'http://localhost:8080/api/v1/insurances';
      const insurance = {

         insuranceIdentification : insurance_identification,
         insuranceAddress      : insurance_address,
         insuranceInsurer     : insurance_insurer,
         //insuranceTimeBuy     : insuranceTimeBuys , //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
         insuranceTimeFirstpay  : insuranceTimeFirstpays  + "T00:00:00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
         productID          : productid,
         memberID        : memberid,
         hospitalID    : hospitalid,
         officerID         : officerid,
      };
      console.log(insurance);
      
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(insurance),
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
          case 'insurance_identification':
            setErrorMessege("error","ระบุเลขประจำตัวประชาชนให้ครบ 13 หลัก");
            return;
          case 'insurance_insurer':
            setErrorMessege("error","ระบุชื่อเป็นภาษาอังกฤษให้ถูกต้อง และตัวแรกเป็นตัวพิมพ์ใหญ่");
            return;
          case 'insurance_address':
            setErrorMessege("error","ระบุที่อยู่ที่ติดต่อได้");
            return;
          default:
            setErrorMessege("error","บันทึกข้อมูลไม่สำเร็จ");
            return;
        }
      }

    return (
      <Page theme={pageTheme.tool}>
        <Header style={HeaderCustom} title={`ระบบการซื้อประกันสุขภาพ`}>
        </Header>
        <Content>
        <ContentHeader title="บันทึกการซื้อประกันสุขภาพ">

     </ContentHeader>
          <Container maxWidth="sm">
            <Grid container spacing={3}>
              <Grid item xs={12}></Grid>
              <Grid item xs={3}>
                <div className={classes.paper}>ผลิตภัณฑ์</div>
              </Grid>
              <Grid item xs={9}>
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
                <div className={classes.paper}>สมาชิกระบบประกันสุขภาพ</div>
              </Grid>
              <Grid item xs={9}>
                  <FormControl
                    className={classes.formControl }
                    variant="outlined"
                  >
                  <TextField
                    id="user"
                    label="ชื่อสมาชิก"
                    type="string"
                    size="medium"
                    variant="outlined"
                    value={members.filter((filter: EntMember) => filter.id == memberid).map((item: EntMember) => `${item.memberName}`)}
                    style={{ width: 300  }}
                  />
                  </FormControl>
              </Grid>

              <Grid item xs={3}>
              <div className={classes.paper}> ราคา </div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl} disabled>
                <InputLabel> </InputLabel>
                <Select
                    name="product"
                    value={productid || ''} // (undefined || '') = ''
                    onChange={product_id_handleChange}
                    startAdornment={<InputAdornment position="start">฿</InputAdornment>}
                  >
                  {products.map((item: EntProduct) => (
                    <MenuItem value={item.id}>{item.productPrice}</MenuItem>
                  ))
                  }
                </Select>
              </FormControl>
            </Grid>
  
              <Grid item xs={3}>
                <div className={classes.paper}>โรงพยาบาล</div>
              </Grid>
              <Grid item xs={9}>
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
                <div className={classes.paper}>พนักงานบริษัทประกันสุขภาพที่แนะนำ</div>
              </Grid>
              <Grid item xs={9}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>เลือกพนักงานบริษัทประกันสุขภาพ</InputLabel>
                  <Select
                    name="officer"
                    value={officerid || ''} // (undefined || '') = ''
                    onChange={officer_id_handleChange}
                  >
                    {officers.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.officerName}
                        </MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เลขประจำตัวประชาชน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  error = {insurance_identificationerror ? true : false}
                  id="outlined-number"
                  name="insurance_identification"
                  multiline
                  rows={4}
                  value={insurance_identification}
                  type="string"
                  onChange={insuranceIdentification_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                  helperText ={insurance_identificationerror}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ข้อมูลของผู้รับผลประโยชน์</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  error = {insurance_insurererror ? true : false}
                  id="outlined-number"
                  name="insurance_insurer"
                  multiline
                  rows={4}
                  value={insurance_insurer}
                  type="string"
                  onChange={insuranceInsurer_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                  helperText ={insurance_insurererror}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ที่อยู่ของผู้ซื้อ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  error = {insurance_addressserror ? true : false}
                  id="outlined-number"
                  name="insurance_address"
                  multiline
                  rows={4}
                  value={insurance_address}
                  type="string"
                  onChange={insuranceAddress_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                  helperText ={insurance_addressserror}
                />
              </FormControl>
            </Grid>

  
              <Grid item xs={3}>
                <div className={classes.paper}>วันที่ต้องการจ่ายงวดแรก</div>
              </Grid>
              <Grid item xs={9}>
                <form className={classes.container} noValidate>
                  <TextField
                    label="เลือกเวลา"
                    name="insurance_timefirstpay"
                    type="date"
                    value={insuranceTimeFirstpays || ''} // (undefined || '') = ''
                    className={classes.textField}
                    InputLabelProps={{
                      shrink: true,
                    }}
                    onChange={insuranceTimeFirstpay_handleChange}
                  />
                </form>
              </Grid>
  
              <Grid item xs={4}></Grid>
              <Grid item xs={8}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={() => {
                  CreateInsurance();
                }}
              >
                บันทึกการซื้อประกันสุขภาพ
              </Button>
              </Grid>
            </Grid>
          </Container>
        </Content>
      </Page>
    );
  }