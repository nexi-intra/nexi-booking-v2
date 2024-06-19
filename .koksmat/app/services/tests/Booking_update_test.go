    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
    //generator:  noma4.1
    package tests
    import (
        "testing"
        "github.com/magicbutton/nexi-booking/services/endpoints/booking"
                    "github.com/magicbutton/nexi-booking/services/models/bookingmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestBookingupdate(t *testing.T) {
                                // transformer v1
            object := bookingmodel.Booking{}
         
            result,err := booking.BookingUpdate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    