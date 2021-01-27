// MIT License
//
// (C) Copyright [2020-2021] Hewlett Packard Enterprise Development LP
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package hms_s3

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type s3_TS struct {
    suite.Suite
}

func (suite *s3_TS) Test_LoadConnectionFromEnvVars_happypath() {
    os.Setenv("S3_ACCESS_KEY", "s3-access-key")
    os.Setenv("S3_SECRET_KEY", "s3-secret-key")
    os.Setenv("S3_ENDPOINT", "http://s3:9000")
    os.Setenv("S3_BUCKET", "tester")
    os.Setenv("S3_REGION", "")

    expectedConnectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://s3:9000",
        Bucket:    "tester",
        Region:    "default",
    }

    returnedConnectionInfo, err := LoadConnectionInfoFromEnvVars()

    suite.NoError(err)
    suite.Equal(expectedConnectionInfo, returnedConnectionInfo)
}

func (suite *s3_TS) Test_LoadConnectionFromEnvVars_missingendpoint() {
    os.Setenv("S3_ACCESS_KEY", "s3-access-key")
    os.Setenv("S3_SECRET_KEY", "s3-secret-key")
    os.Setenv("S3_ENDPOINT", "")
    os.Setenv("S3_BUCKET", "tester")
    os.Setenv("S3_REGION", "")

    expectedConnectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://s3:9000",
        Bucket:    "tester",
        Region:    "default",
    }

    returnedConnectionInfo, err := LoadConnectionInfoFromEnvVars()

    suite.Error(err)
    suite.NotEqual(expectedConnectionInfo, returnedConnectionInfo)
}

func (suite *s3_TS) Test_LoadConnectionFromEnvVars_missingaccess() {
    os.Setenv("S3_ACCESS_KEY", "")
    os.Setenv("S3_SECRET_KEY", "s3-secret-key")
    os.Setenv("S3_ENDPOINT", "http://s3:9000")
    os.Setenv("S3_BUCKET", "tester")
    os.Setenv("S3_REGION", "")

    expectedConnectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://s3:9000",
        Bucket:    "tester",
        Region:    "default",
    }

    returnedConnectionInfo, err := LoadConnectionInfoFromEnvVars()

    suite.Error(err)
    suite.NotEqual(expectedConnectionInfo, returnedConnectionInfo)
}

func (suite *s3_TS) Test_LoadConnectionFromEnvVars_missingsecret() {
    os.Setenv("S3_ACCESS_KEY", "s3-access-key")
    os.Setenv("S3_SECRET_KEY", "")
    os.Setenv("S3_ENDPOINT", "http://s3:9000")
    os.Setenv("S3_BUCKET", "tester")
    os.Setenv("S3_REGION", "")

    expectedConnectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://s3:9000",
        Bucket:    "tester",
        Region:    "default",
    }

    returnedConnectionInfo, err := LoadConnectionInfoFromEnvVars()

    suite.Error(err)
    suite.NotEqual(expectedConnectionInfo, returnedConnectionInfo)
}

func (suite *s3_TS) Test_LoadConnectionFromEnvVars_notequal() {
    os.Setenv("S3_ACCESS_KEY", "garbage")
    os.Setenv("S3_SECRET_KEY", "s3-secret-key")
    os.Setenv("S3_ENDPOINT", "http://s3:9000")
    os.Setenv("S3_BUCKET", "tester")
    os.Setenv("S3_REGION", "")

    expectedConnectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://s3:9000",
        Bucket:    "tester",
        Region:    "default",
    }

    returnedConnectionInfo, err := LoadConnectionInfoFromEnvVars()

    suite.NoError(err)
    suite.NotEqual(expectedConnectionInfo, returnedConnectionInfo)
}


/*
Other test cases:

    test good ping
    change the bucket => bad ping

 */

func (suite *s3_TS) Test_NewConnection_happypath() {

    expectedConnectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://s3:9000",
        Bucket:    "tester",
        Region:    "default",
    }


    returnedConnectionInfo := NewConnectionInfo(expectedConnectionInfo.AccessKey, expectedConnectionInfo.SecretKey, expectedConnectionInfo.Endpoint, expectedConnectionInfo.Bucket, expectedConnectionInfo.Region)
    suite.Equal(expectedConnectionInfo, returnedConnectionInfo)
}

func (suite *s3_TS) Test_NewConnection_notequal() {

    expectedConnectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://s3:9000",
        Bucket:    "tester",
        Region:    "default",
    }


    returnedConnectionInfo := NewConnectionInfo(expectedConnectionInfo.SecretKey, expectedConnectionInfo.SecretKey, expectedConnectionInfo.Endpoint, expectedConnectionInfo.Bucket, expectedConnectionInfo.Region)
    suite.NotEqual(expectedConnectionInfo, returnedConnectionInfo)
}

func (suite *s3_TS) Test_NewClient_happypath() {

    expectedConnectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://s3:9000",
        Bucket:    "tester",
        Region:    "default",
    }

    _, err := NewS3Client(expectedConnectionInfo, nil)

    suite.NoError(err)
}

func (suite *s3_TS) Test_NewClient_badEndpoint() {

    expectedConnectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://notfound:9000",
        Bucket:    "tester",
        Region:    "default",
    }

    client, err := NewS3Client(expectedConnectionInfo, nil)
    suite.NoError(err)

    err = client.PingBucket()
    suite.Error(err)
}

func (suite *s3_TS) Test_PingBucket_happypath() {

    expectedConnectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://s3:9000",
        Bucket:    "tester",
        Region:    "default",
    }

    client, err := NewS3Client(expectedConnectionInfo, nil)
    suite.NoError(err)

    err = client.PingBucket()
    suite.NoError(err)
}

func (suite *s3_TS) Test_PingBucket_badBucket() {

    expectedConnectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://s3:9000",
        Bucket:    "tester",
        Region:    "default",
    }

    client, err := NewS3Client(expectedConnectionInfo, nil)

    suite.NoError(err)
    client.SetBucket("garbage")

    err = client.PingBucket()
    suite.Error(err)
}

func (suite *s3_TS) Test_ConnectioninfoValidate_HappyPath() {
    connectionInfo := ConnectionInfo{
        AccessKey: "s3-access-key",
        SecretKey: "s3-secret-key",
        Endpoint:  "http://s3:9000",
        Bucket:    "tester",
        Region:    "default",
    }

    err := connectionInfo.Validate()
    suite.NoError(err)
}

func (suite *s3_TS) Test_ConnectioninfoValidate_Invalid() {
    tests := []struct {
        connectionInfo ConnectionInfo
        expectedError error

    }{{
        // Empty Access Key
        connectionInfo: ConnectionInfo{
            SecretKey: "s3-secret-key",
            Endpoint:  "http://s3:9000",
            Bucket:    "tester",
            Region:    "default",    
        },
        expectedError: errors.New("s3 access key is empty"),
    }, {
        // Empty Secret Key
        connectionInfo: ConnectionInfo{
            AccessKey: "s3-access-key",
            Endpoint:  "http://s3:9000",
            Bucket:    "tester",
            Region:    "default",    
        },
        expectedError: errors.New("s3 secret key is empty"),
    }, {
        // Empty Endpoint
        connectionInfo: ConnectionInfo{
            AccessKey: "s3-access-key",
            SecretKey: "s3-secret-key",
            Bucket:    "tester",
            Region:    "default",    
        },
        expectedError: errors.New("s3 endpoint is empty"),
    }, {
        // Empty Bucket
        connectionInfo: ConnectionInfo{
            AccessKey: "s3-access-key",
            SecretKey: "s3-secret-key",
            Endpoint:  "http://s3:9000",
            Region:    "default",    
        },
        expectedError: errors.New("s3 bucket is empty"),
    }, {
        // Empty Region
        connectionInfo: ConnectionInfo{
            AccessKey: "s3-access-key",
            SecretKey: "s3-secret-key",
            Endpoint:  "http://s3:9000",
            Bucket:    "tester",
        },
        expectedError: errors.New("s3 region is empty"),
    }}

    for _, test := range tests {
		err := test.connectionInfo.Validate()
		suite.Equal(test.expectedError, err)
	}
}


func TestS3Suite(t *testing.T) {
    suite.Run(t, new(s3_TS))
}
