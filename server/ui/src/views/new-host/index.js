import React from 'react'
import styled from 'styled-components'

import { useFormState } from 'react-use-form-state'

import Flex from '_components/flex'
import Box from '_components/box'
import Link from '_components/link'
import Text from '_components/text'
import Button from '_components/button'
import TextInput from '_components/inputs/text-input'
import { Dragon as Spinner } from '_components/spinner'

import { useMutation } from 'react-apollo'
import { CREATE_HOST } from '_graphql/actions'
import { navigate } from '@reach/router'
import toast from '_components/toast'

const Container = styled(Flex)`
  flex-direction: column;
`

const NewHost = () => {
  const [formState, { text }] = useFormState()

  const onHostCreated = () => {
    toast.success('Host created')
    navigate('/hosts')
  }

  const onHostCreationError = () => {
    toast.error('Error creating host')
    navigate('/hosts')
  }

  const [createHost, { loading }] = useMutation(CREATE_HOST, {
    variables: formState.values,
    onCompleted: onHostCreated,
    onError: onHostCreationError,
  })

  const onSave = () => {
    createHost()
  }

  return (
    <Container>
      <Text textStyle="title" mb={4}>
        New host
      </Text>
      {loading ? (
        <Spinner />
      ) : (
        <Box flexDirection="column">
          <Text my={3}>Name</Text>
          <TextInput required {...text('name')} placeholder="new-host-1" mb={2} />
          <Text my={3}>Address</Text>
          <TextInput required {...text('address')} placeholder="10.0.8.0/24" mb={2} />
          <Text my={3}>Advertise address</Text>
          <TextInput required {...text('advertiseAddress')} placeholder="wg.domain.com" mb={2} />
          <Text my={3}>Listen port</Text>
          <TextInput required {...text('listenPort')} placeholder="51820" mb={2} />
          <Button width="100%" borderRadius={3} mt={3} mb={4} onClick={onSave}>
            Save
          </Button>
        </Box>
      )}

      <Box justifyContent="center" gridColumn="4 / span 6">
        <Link color="primary" to="/hosts">
          Cancel
        </Link>
      </Box>
    </Container>
  )
}

export default NewHost